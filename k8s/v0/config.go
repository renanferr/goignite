package gik8s

import (
	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	configRoot     = "gi.k8s"
	kubeConfigPath = configRoot + ".kubeConfigPath"
)

func init() {
	giconfig.Add(kubeConfigPath, "~/.kube/config", "defines kubeconfig request")
}
