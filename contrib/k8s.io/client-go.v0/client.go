package client

import (
	"context"

	"github.com/b2wdigital/goignite/v2/core/log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func NewClient(ctx context.Context, options *Options) *kubernetes.Clientset {

	logger := log.FromContext(ctx).
		WithField("context", options.Context).
		WithField("kubeConfigPath", options.KubeConfigPath)

	logger.Tracef("creating k8s client")

	config, err := fromKubeConfig(options.Context, options.KubeConfigPath)
	if err != nil {
		logger.Error(err.Error())
		return nil
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Error(err.Error())
		return nil
	}

	return client
}

func fromKubeConfig(context string, kubeConfigPath string) (*rest.Config, error) {
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeConfigPath},
		&clientcmd.ConfigOverrides{
			CurrentContext: context,
		}).ClientConfig()
}
