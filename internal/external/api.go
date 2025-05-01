package external

import (
	"context"
	"fmt"
	"github.com/alserok/g8s/internal/service/models"
)

type KubernetesClient interface {
	CreateDeployment(ctx context.Context, dep models.Deployment) error
	CreatePersistentVolumeClaim(ctx context.Context, pvc models.PersistentVolumeClaim) error
	CreateService(ctx context.Context, srvc models.Service) error

	DeleteDeployment(ctx context.Context, namespace, name string) error

	UpdateDeployment(ctx context.Context, dep models.Deployment) error
	UpdatePersistentVolumeClaim(ctx context.Context, pvc models.PersistentVolumeClaim) error
	UpdateService(ctx context.Context, srvc models.Service) error

	ListDeployments(ctx context.Context, namespace string) ([]models.Deployment, error)
}

type AIClient interface {
	Prompt(ctx context.Context, prompt fmt.Stringer) (string, error)
}
