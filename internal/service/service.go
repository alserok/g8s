package service

import (
	"context"
	"fmt"
	"github.com/alserok/g8s/internal/external"
	"github.com/alserok/g8s/internal/metrics"
	"github.com/alserok/g8s/internal/service/models"
	"github.com/alserok/g8s/internal/utils/errors"
	"github.com/alserok/g8s/internal/utils/helpers"
)

type Service interface {
	List(ctx context.Context, req models.List) (any, error)
	Create(ctx context.Context, req models.Create) error
	Delete(ctx context.Context, req models.Delete) error
	Update(ctx context.Context, req models.Update) error
}

func New(k8sClient external.KubernetesClient, aiClient external.AIClient, metr metrics.Metrics) Service {
	return &service{
		k8sClient: k8sClient,
		aiClient:  aiClient,
	}
}

type service struct {
	k8sClient external.KubernetesClient
	aiClient  external.AIClient
}

func (s service) List(ctx context.Context, req models.List) (any, error) {
	switch req.Type {
	case models.TypeDeployment:
		deps, err := s.k8sClient.ListDeployments(ctx, req.Namespace)
		if err != nil {
			return nil, fmt.Errorf("faild to list deployments: %w", err)
		}

		return deps, err
	default:
		return nil, errors.New("unknown type", errors.ErrBadRequest)
	}
}

func (s service) Create(ctx context.Context, req models.Create) error {
	schema, err := s.aiClient.Prompt(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to prompt ai: %w", err)
	}

	switch req.Type {
	case models.TypeDB:
		schemas, err := helpers.ParseSchema(ctx, schema)
		if err != nil {
			return fmt.Errorf("failed to parse schema: %w", err)
		}

		if dep, ok := schemas[models.TypeDeployment]; ok {
			if err = s.k8sClient.CreateDeployment(ctx, dep.(models.Deployment)); err != nil {
				return fmt.Errorf("failed to create deployment: %w", err)
			}
		}

		if dep, ok := schemas[models.TypePersistentVolumeClaim]; ok {
			if err = s.k8sClient.CreatePersistentVolumeClaim(ctx, dep.(models.PersistentVolumeClaim)); err != nil {
				return fmt.Errorf("failed to create persistent volume claim: %w", err)
			}
		}

		if srvc, ok := schemas[models.TypeService]; ok {
			if err = s.k8sClient.CreateService(ctx, srvc.(models.Service)); err != nil {
				return fmt.Errorf("failed to create persistent volume claim: %w", err)
			}
		}
	case models.TypeService:
	default:
		return errors.New("unknown type", errors.ErrBadRequest)
	}

	return nil
}

func (s service) Delete(ctx context.Context, req models.Delete) error {
	if err := s.k8sClient.DeleteDeployment(ctx, req.Namespace, req.Name); err != nil {
		return fmt.Errorf("failed to delete deployment: %w", err)
	}

	return nil
}

func (s service) Update(ctx context.Context, req models.Update) error {
	schema, err := s.aiClient.Prompt(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to prompt ai: %w", err)
	}

	switch req.Type {
	case models.TypeDB:
		schemas, err := helpers.ParseSchema(ctx, schema)
		if err != nil {
			return fmt.Errorf("failed to parse schema: %w", err)
		}

		if dep, ok := schemas[models.TypeDeployment]; ok {
			if err = s.k8sClient.UpdateDeployment(ctx, dep.(models.Deployment)); err != nil {
				return fmt.Errorf("failed to update deployment: %w", err)
			}
		}

		if dep, ok := schemas[models.TypePersistentVolumeClaim]; ok {
			if err = s.k8sClient.UpdatePersistentVolumeClaim(ctx, dep.(models.PersistentVolumeClaim)); err != nil {
				return fmt.Errorf("failed to update persistent volume claim: %w", err)
			}
		}

		if srvc, ok := schemas[models.TypeService]; ok {
			if err = s.k8sClient.UpdateService(ctx, srvc.(models.Service)); err != nil {
				return fmt.Errorf("failed to update persistent volume claim: %w", err)
			}
		}
	case models.TypeService:
	default:
		return errors.New("unknown type", errors.ErrBadRequest)
	}

	return nil
}
