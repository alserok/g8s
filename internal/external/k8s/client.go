package k8s

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/alserok/g8s/internal/external"
)

func NewClient(kubeConfig string) external.KubernetesClient {
	cl, err := kubernetes.NewForConfig(&rest.Config{})
	if err != nil {
		panic("failed to init client: " + err.Error())
	}

	return &client{cl: cl}
}

type client struct {
	cl *kubernetes.Clientset
}
