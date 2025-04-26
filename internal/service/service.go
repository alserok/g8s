package service

import (
	"context"
	"github.com/alserok/g8s/internal/external"
	"github.com/alserok/g8s/internal/service/domains"
	"github.com/alserok/g8s/internal/service/models"
)

type Service interface {
	GetDeployments(ctx context.Context, namespace string) ([]models.Deployment, error)
}

func New(k8sClient external.KubernetesClient) Service {
	return &service{}
}

type service struct {
	domains.Deployments
}
