package v1

import (
	"net/http"

	"github.com/alserok/g8s/internal/service"
)

type Handler struct {
	Service service.Service
}

func (h *Handler) GetPods(w http.ResponseWriter, r *http.Request) error {
	return nil
}
