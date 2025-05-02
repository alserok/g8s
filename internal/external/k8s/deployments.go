package k8s

import (
	"context"
	"github.com/alserok/g8s/internal/service/models"
	"github.com/alserok/g8s/internal/utils/errors"
	"github.com/alserok/g8s/internal/utils/logger"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *client) CreateDeployment(ctx context.Context, dep models.Deployment) error {
	cl := c.cl.AppsV1().Deployments(dep.Namespace)

	d := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: dep.ObjectMeta.Name,
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
	log := logger.ExtractContext(ctx)

	log.Debug("listing deployments", logger.WithArg("namespace", namespace))

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

	log.Debug("deployments list", logger.WithArg("namespace", namespace), logger.WithArg("deployments", deps))

	return deps, nil
}

func (c *client) DeleteDeployment(ctx context.Context, namespace, name string) error {
	cl := c.cl.AppsV1().Deployments(namespace)

	if err := cl.Delete(ctx, name, metav1.DeleteOptions{}); err != nil {
		return errors.New(err.Error(), errors.ErrInternal)
	}

	return nil
}

func (c *client) UpdateDeployment(ctx context.Context, dep models.Deployment) error {
	cl := c.cl.AppsV1().Deployments(dep.Namespace)

	d := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: dep.ObjectMeta.Name,
		},
		Spec:   appsv1.DeploymentSpec{},
		Status: appsv1.DeploymentStatus{},
	}

	if _, err := cl.Update(ctx, d, metav1.UpdateOptions{}); err != nil {
		return errors.New(err.Error(), errors.ErrInternal)
	}

	return nil
}

func (c *client) UpdatePersistentVolumeClaim(ctx context.Context, pvc models.PersistentVolumeClaim) error {
	cl := c.cl.AppsV1().Deployments("")

	d := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "",
		},
		Spec:   appsv1.DeploymentSpec{},
		Status: appsv1.DeploymentStatus{},
	}

	if _, err := cl.Update(ctx, d, metav1.UpdateOptions{}); err != nil {
		return errors.New(err.Error(), errors.ErrInternal)
	}

	return nil
}

func (c *client) UpdateService(ctx context.Context, srvc models.Service) error {
	//TODO implement me
	panic("implement me")
}
