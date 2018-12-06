package manager

import (
	"encoding/json"
	"net"

	"github.com/ohmk/k8s-edge/api"
	clientset "github.com/ohmk/k8s-edge/pkg/client/clientset/versioned"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type syncService struct {
	kubeclientset kubernetes.Interface
	edgeclientset clientset.Interface

	grpcServer *grpc.Server
	lis        net.Listener
}

func newSyncService(masterURL, path, l string) (*syncService, error) {
	config, err := clientcmd.BuildConfigFromFlags(masterURL, path)
	if err != nil {
		log.Fatalf("Failed to build config from: %s, %s", masterURL, path)
		return nil, err
	}
	kubeclientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Failed to build config from: %s, %s", masterURL, path)
		return nil, err
	}
	edgeclientset, err := clientset.NewForConfig(config)
	if err != nil {
		log.Fatalf("Failed to build config from: %s, %s", masterURL, path)
		return nil, err
	}

	lis, err := net.Listen("tcp", l)
	if err != nil {
		log.Fatalf("Failed to listen port: %s", l)
		return nil, err
	}

	s := &syncService{
		kubeclientset: kubeclientset,
		edgeclientset: edgeclientset,
		grpcServer:    grpc.NewServer(),
		lis:           lis,
	}
	api.RegisterSyncServiceAPIServer(s.grpcServer, s)

	return s, nil
}

func (s *syncService) Run(errCh chan<- error) {
	log.Info("Starting sync service api server")
	s.grpcServer.Serve(s.lis)
}

func (s *syncService) GetEdgeNode(ctx context.Context, in *api.GetEdgeNodeRequest) (*api.GetEdgeNodeReply, error) {
	log.Infof("GetEdgeNode() is called from node '%s'", in.NodeName)

	e, err := s.edgeclientset.EdgeV1alpha1().EdgeNodes("default").Get(in.NodeName, metav1.GetOptions{})
	if err != nil {
		log.Warn(err)
		return &api.GetEdgeNodeReply{}, nil
	}
	// TODO: should use UpdateStatus
	e.Status.LastSyncedAt = metav1.Now()
	_, err = s.edgeclientset.EdgeV1alpha1().EdgeNodes("default").Update(e)

	if e.Status.Phase == "Running" {
		log.Info("There is nothing to sync")
		return &api.GetEdgeNodeReply{
			EdgeNode: &api.EdgeNode{
				Name:   in.NodeName,
				Synced: true,
				Pods:   nil,
			},
		}, nil
	}

	pods := make([]*api.Pod, 0)
	for _, pod := range e.Spec.Pods {
		jsonBytes, err := json.Marshal(pod)
		if err != nil {
			log.Warnf("Failed to serialize pod '%s' to json", pod.ObjectMeta.Name)
			continue
		}
		p := &api.Pod{
			Key:   pod.ObjectMeta.Name,
			Value: jsonBytes,
		}
		pods = append(pods, p)
	}

	return &api.GetEdgeNodeReply{
		EdgeNode: &api.EdgeNode{
			Name:   in.NodeName,
			Synced: false,
			Pods:   pods,
		},
	}, nil
}
