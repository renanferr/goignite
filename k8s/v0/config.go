package gik8s

import (
	giconfig "github.com/b2wdigital/goignite/config"

	"log"
)

const (
	configRoot                     = "gi.k8s"
	kubeConfigPath                 = configRoot + ".kubeConfigPath"
)

func init() {
	log.Println("getting configurations for k8s")

	giconfig.Add(kubeConfigPath, "~/.kube/config", "defines kubeconfig request")
}
