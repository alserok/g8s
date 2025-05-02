package k8s

import (
	"github.com/alserok/g8s/internal/external"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func NewClient() external.KubernetesClient {
	cfg, err := rest.InClusterConfig()
	if err != nil {
		panic("failed to init k8s client config: " + err.Error())
	}

	cl, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		panic("failed to init client: " + err.Error())
	}

	return &client{cl: cl}
}

type client struct {
	cl *kubernetes.Clientset
}
