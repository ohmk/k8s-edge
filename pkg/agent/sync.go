package agent

import (
	"context"
	"os"
	"reflect"
	"time"

	"github.com/ohmk/k8s-edge/api"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// TODO
var filePath map[string]bool = map[string]bool{}

type syncAgent struct {
	name            string
	podManifestPath string
	serverURL       string
	timeout         int
	interval        int
}

func newSyncAgent(nodeName, podManifestPath, serverURL string, timeout, interval int) (*syncAgent, error) {
	return &syncAgent{
		name:            nodeName,
		podManifestPath: podManifestPath,
		serverURL:       serverURL,
		timeout:         timeout,
		interval:        interval,
	}, nil
}

func (s *syncAgent) Run(errCh chan<- error) {
	for {
		doneCh := make(chan struct{})
		ctx, _ := context.WithTimeout(context.Background(), time.Duration(s.timeout)*time.Second)

		go s.syncEdgeSpec(ctx, doneCh)

		select {
		case <-doneCh:
			log.Info("Syncing edge spec was done")
			time.Sleep(time.Duration(s.interval) * time.Second)
		case <-ctx.Done():
			log.Info("Syncing edge spec was timeout")
			time.Sleep(time.Duration(s.interval) * time.Second)
		}
	}
}

func (s *syncAgent) syncEdgeSpec(ctx context.Context, doneCh chan struct{}) {
	log.Info("Starting sync edge spec")
	defer close(doneCh)

	conn, err := grpc.Dial(s.serverURL, grpc.WithInsecure())
	if err != nil {
		// TODO: notify error to the parent process
		log.Warn("Failed to connect sync server:", err)
	}
	defer conn.Close()

	c := api.NewSyncServiceAPIClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &api.GetEdgeNodeRequest{
		NodeName: s.name,
	}
	res, err := c.GetEdgeNode(ctx, req)
	if err != nil {
		// TODO: notify error to the parent process
		log.Warn("Failed to get node info: ", err)
		return
	}

	// EdgeNode resource has been deleted
	if reflect.DeepEqual(res, &api.GetEdgeNodeReply{}) {
		log.Infof("Edgenode '%s' resouce is not found", s.name)

		for path, _ := range filePath {
			err := os.Remove(path)
			if err != nil {
				log.Warnf("Failed to remove %s: %v", path, err)
			}
		}
		filePath = map[string]bool{}
		return
	}

	// EdgeNode resource has already been synced
	if res.EdgeNode.Synced {
		log.Info("Already synced")
		return
	}

	// Try to match the actual state to the desired state
	for path, _ := range filePath {
		filePath[path] = false
	}
	for _, pod := range res.EdgeNode.Pods {
		path := s.podManifestPath + "/" + pod.Key + ".json"
		_, ok := filePath[path]
		if ok {
			log.Infof("%v has already been created", path)
			filePath[path] = true
			continue
		}

		file, err := os.Create(path)
		if err != nil {
			log.Infof("Failed to create file '%s': %v", path, err)
		}
		defer file.Close()
		file.Write(pod.Value)
		filePath[path] = true
	}
	for path, flag := range filePath {
		if !flag {
			err := os.Remove(path)
			if err != nil {
				log.Warnf("Failed to remove %s: %v", path, err)
			}
			delete(filePath, path)
		}
	}
}
