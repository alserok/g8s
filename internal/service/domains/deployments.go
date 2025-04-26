package domains

import (
	"context"
	"fmt"
	"github.com/alserok/g8s/internal/external"
	"github.com/alserok/g8s/internal/service/models"
)

type Deployments struct {
	K8sCl external.KubernetesClient
}

func (d Deployments) ListDeployments(ctx context.Context, namespace string) ([]models.Deployment, error) {
	deps, err := d.K8sCl.ListDeployments(ctx, namespace)
	if err != nil {
		return nil, fmt.Errorf("on ListDeployments: %w", err)
	}

	return deps, nil
}

func (d Deployments) CreateDeployment(ctx context.Context, dep models.Deployment) error {
	err := d.K8sCl.CreateDeployment(ctx, dep)
	if err != nil {
		return fmt.Errorf("on CreateDeployment: %w", err)
	}

	return nil
}
