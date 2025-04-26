package service

import (
	"context"
	"github.com/alserok/g8s/internal/external"
	"github.com/alserok/g8s/internal/service/domains"
	"github.com/alserok/g8s/internal/service/models"
)

type Service interface {
	ListDeployments(ctx context.Context, namespace string) ([]models.Deployment, error)
	CreateDeployment(ctx context.Context, dep models.Deployment) error
}

func New(k8sClient external.KubernetesClient) Service {
	return &service{
		domains.Deployments{K8sCl: k8sClient},
	}
}

type service struct {
	domains.Deployments
}
