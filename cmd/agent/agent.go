package main

import (
	"os"

	"github.com/ohmk/k8s-edge/pkg/agent"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

// TODO
var (
	defaultNodeName, _ = os.Hostname()
	nodeName           = kingpin.Flag(
		"node-name",
		"The name of this node.",
	).Default(defaultNodeName).String()
	podManifestPath = kingpin.Flag(
		"pod-manifest-path",
		"Path to the directory containing static pod files to run.",
	).Default("/etc/kubernetes/manifests").String()
	syncServerURL = kingpin.Flag(
		"sync-server",
		"The address of the sync server.",
	).Default("localhost:50051").String()
	syncTimeout = kingpin.Flag(
		"sync-timeout",
		"Connection timeout to sync server.",
	).Default("5").Int()
	syncInterval = kingpin.Flag(
		"sync-interval",
		"Connection interval to sync server.",
	).Default("10").Int()
)

func main() {
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	a, err := agent.New(*nodeName, *podManifestPath, *syncServerURL, *syncTimeout, *syncInterval)
	if err != nil {
		log.Fatal("Failed to initialize edge-agent:", err)
	}

	log.Info("Starting edge agent")
	a.Run()
}
