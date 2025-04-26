package external

import (
	"context"
	"github.com/alserok/g8s/internal/service/models"
)

type KubernetesClient interface {
	CreateDeployment(ctx context.Context, depl models.Deployment) error
	ListDeployments(ctx context.Context, namespace string) ([]models.Deployment, error)
}

type AIClient interface {
	Prompt(ctx context.Context, prompt string) (string, error)
}
