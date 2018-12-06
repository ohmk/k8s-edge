package agent

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

type agent struct {
	sync *syncAgent
}

func New(nodeName, podManifestPath, syncServerURL string, syncTimeout, syncInterval int) (*agent, error) {
	s, _ := newSyncAgent(nodeName, podManifestPath, syncServerURL, syncTimeout, syncInterval)

	return &agent{
		sync: s,
	}, nil
}

func (a *agent) Run() {
	errCh := make(chan error, 1)
	sigCh := setupSignalHandler()

	log.Info("Starting sync agent")
	go a.sync.Run(errCh)

	select {
	case <-sigCh:
		log.Warn("Terminating edge agent")
		// TODO: shutdown gracefully
		return
	case err := <-errCh:
		log.Fatal("Terminating edge agent:", err)
		return
	}
}

func setupSignalHandler() (stopCh <-chan struct{}) {
	stop := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, []os.Signal{os.Interrupt, syscall.SIGTERM}...)
	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}
