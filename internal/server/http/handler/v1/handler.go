package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alserok/g8s/internal/service"
	"github.com/alserok/g8s/internal/service/models"
	"github.com/alserok/g8s/internal/utils/errors"
)

type Handler struct {
	Service service.Service
}

func (h *Handler) GetDeployments(w http.ResponseWriter, r *http.Request) error {
	deps, err := h.Service.ListDeployments(r.Context(), r.PathValue("namespace"))
	if err != nil {
		return fmt.Errorf("on ListDeployments: %w", err)
	}

	_ = json.NewEncoder(w).Encode(deps)

	return nil
}

func (h *Handler) CreateDeployment(w http.ResponseWriter, r *http.Request) error {
	var dep models.Deployment
	if err := json.NewDecoder(r.Body).Decode(&dep); err != nil {
		return errors.New(err.Error(), errors.ErrBadRequest)
	}

	if err := h.Service.CreateDeployment(r.Context(), dep); err != nil {
		return fmt.Errorf("on CreateDeployment: %w", err)
	}

	return nil
}
