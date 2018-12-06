package manager

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

type manager struct {
	sync *syncService
	edge *edgeController
}

func New(masterURL, path, lis string) (*manager, error) {
	s, err := newSyncService(masterURL, path, lis)
	if err != nil {
		log.Fatal("Failed to initialize sync service")
		return nil, err
	}

	e, err := newEdgeController(masterURL, path)
	if err != nil {
		log.Fatal("Failed to initialize edge controller")
		return nil, err
	}

	return &manager{
		sync: s,
		edge: e,
	}, nil
}

func (m *manager) Run() {
	errCh := make(chan error, 1)
	sigCh := setupSignalHandler()

	log.Info("Starting sync service")
	go m.sync.Run(errCh)

	log.Info("Starting edge controller")
	go m.edge.Run(sigCh, errCh)

	select {
	case <-sigCh:
		log.Warn("Terminating edge manager")
		// TODO: shutdown gracefully
		return
	case err := <-errCh:
		log.Fatal("Terminated edge manager:", err)
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
