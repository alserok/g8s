package domains

import (
	"context"
	"github.com/alserok/g8s/internal/service/models"
)

type Deployments struct {
}

func (d Deployments) GetDeployments(ctx context.Context, namespace string) ([]models.Deployment, error) {
	//TODO implement me
	panic("implement me")
}
