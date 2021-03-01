package gik8s

import (
	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	root           = "gi.k8s"
	kubeConfigPath = root + ".kubeConfigPath"
)

func init() {
	giconfig.Add(kubeConfigPath, "~/.kube/config", "defines kubeconfig request")
}
