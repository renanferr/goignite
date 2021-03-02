package gik8s

import (
	"context"

	gilog "github.com/b2wdigital/goignite/log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func NewClient(ctx context.Context, options *Options) *kubernetes.Clientset {

	logger := gilog.FromContext(ctx).
		WithField("context", options.Context).
		WithField("kubeConfigPath", options.KubeConfigPath)

	logger.Infof("creating k8s client")

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
