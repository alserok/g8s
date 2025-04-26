package k8s

import (
	"context"
	"github.com/alserok/g8s/internal/service/models"
	"github.com/alserok/g8s/internal/utils/errors"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *client) CreateDeployment(ctx context.Context, depl models.Deployment) error {
	cl := c.cl.AppsV1().Deployments(depl.Namespace)

	d := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: depl.ObjectMeta.Name,
		},
		Spec:   appsv1.DeploymentSpec{},
		Status: appsv1.DeploymentStatus{},
	}

	_, err := cl.Create(ctx, d, metav1.CreateOptions{})
	if err != nil {
		return errors.New(err.Error(), errors.ErrInternal)
	}

	return nil
}

func (c *client) ListDeployments(ctx context.Context, namespace string) ([]models.Deployment, error) {
	cl := c.cl.AppsV1().Deployments(namespace)

	list, err := cl.List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, errors.New(err.Error(), errors.ErrInternal)
	}

	deps := make([]models.Deployment, 0, len(list.Items))
	for _, dep := range deps {
		deps = append(deps, models.Deployment{
			Namespace:  dep.Namespace,
			ObjectMeta: models.ObjectMeta{},
			Spec:       models.Spec{},
			Template:   models.Template{},
		})
	}

	return deps, nil
}
