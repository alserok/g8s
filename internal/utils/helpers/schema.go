package helpers

import (
	"context"
	"github.com/alserok/g8s/internal/utils/errors"
	"github.com/alserok/g8s/internal/utils/logger"
	"k8s.io/apimachinery/pkg/util/yaml"
	"strings"

	"github.com/alserok/g8s/internal/service/models"
)

func ParseSchema(ctx context.Context, schema string) (map[models.Type]any, error) {
	log := logger.ExtractContext(ctx)

	log.Debug("schema to parse", logger.WithArg("schema", schema))

	parts := strings.Split(schema, "---")

	res := make(map[models.Type]any)

	for _, part := range parts {
		switch {
		case strings.Contains(part, "kind: Service"):
			var srvc models.Service
			if err := yaml.Unmarshal([]byte(part), &srvc); err != nil {
				return nil, errors.New(err.Error(), errors.ErrInternal)
			}

			res[models.TypeService] = srvc
		case strings.Contains(part, "kind: PersistentVolumeClaim"):
			var pvc models.PersistentVolumeClaim
			if err := yaml.Unmarshal([]byte(part), &pvc); err != nil {
				return nil, errors.New(err.Error(), errors.ErrInternal)
			}

			res[models.TypePersistentVolumeClaim] = pvc
		case strings.Contains(part, "kind: Deployment"):
			var dep models.Deployment
			if err := yaml.Unmarshal([]byte(part), &dep); err != nil {
				return nil, errors.New(err.Error(), errors.ErrInternal)
			}

			res[models.TypeDeployment] = dep
		}
	}

	return res, nil
}
