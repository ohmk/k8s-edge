package main

import (
	"os"
	"path/filepath"

	"github.com/ohmk/k8s-edge/pkg/manager"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

// TODO: define struct for configuration
var (
	defaultKubeConfigPath = filepath.Join(os.Getenv("HOME"), ".kube", "config")

	kubeConfigPath = kingpin.Flag(
		"kubeconfig",
		"Absolute path to the kubeconfig file.",
	).Default(defaultKubeConfigPath).String()
	masterURL = kingpin.Flag(
		"master",
		"The address of the Kubernetes API server.",
	).Default("").String()
	listenAddress = kingpin.Flag(
		"sync-server",
		"The address of the sync server.",
	).Default(":50051").String()
)

func main() {
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	m, err := manager.New(*masterURL, *kubeConfigPath, *listenAddress)
	if err != nil {
		log.Fatal("Failed to initialize edge-manager:", err)
	}

	log.Info("Starting edge manager")
	m.Run()
}
