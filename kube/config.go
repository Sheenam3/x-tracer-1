package kube

import (
	"flag"
	"os"
	"path/filepath"
)

type config struct {
	HomeDir    string
	KubeConfig *string
	Frequency  int
	Setup      bool
	AskVersion bool
	AskHelp    bool
	Debug      bool
}

// Check if configuration is initialized
func (c *config) Initialized() bool {
	return c != nil && c.Setup
}

var CONFIG config

// Get configuration
func GetConfig() config {
	if CONFIG.Initialized() {
		return CONFIG
	}

	c := config{}

	// Home directory
	// FIXME replace by HomeDir() // k8s.io/client-go/kubernetes/util
	c.HomeDir = os.Getenv("USERPROFILE")
	if h := os.Getenv("HOME"); h != "" {
		c.HomeDir = h
	}

	// Kubernetes configuration
	if h := c.HomeDir; h != "" {
		c.KubeConfig = flag.String("kubeconfig", filepath.Join(h, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		c.KubeConfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	// Refreshing frequency
	flag.IntVar(&c.Frequency, "frequency", 3, "refreshing frequency in seconds (default: 5)")

	// CLI Asks
	flag.BoolVar(&c.AskVersion, "version", false, "get current version")
	flag.BoolVar(&c.AskHelp, "help", false, "get help")
	flag.BoolVar(&c.Debug, "kind", false, "for kind env.")
	flag.Parse()

	c.Setup = true
	CONFIG = c

	return c
}
