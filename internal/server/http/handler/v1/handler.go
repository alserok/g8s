package v1

import (
	"encoding/json"
	"fmt"
	_ "github.com/swaggo/http-swagger"
	"net/http"

	"github.com/alserok/g8s/internal/service"
	"github.com/alserok/g8s/internal/service/models"
	"github.com/alserok/g8s/internal/utils/errors"
)

type Handler struct {
	Service service.Service
}

// List godoc
// @Summary      List
// @Description  lists entities
// @Tags         v1
// @Accept       json
// @Produce      json
// @Param        namespace   path      string  true  "namespace"
// @Success      200  {int}  0
// @Failure      400  {int}  0
// @Failure      404  {int}  0
// @Failure      500  {int}  0
// @Router       /v1/list/{namespace} [get]
func (h *Handler) List(w http.ResponseWriter, r *http.Request) error {
	var req models.List

	req.Namespace = r.PathValue("namespace")

	//t, err := strconv.Atoi(r.PathValue(""))
	//if err != nil {
	//	return errors.New(err.Error(), errors.ErrBadRequest)
	//}
	//req.Type = models.Type(t)

	deps, err := h.Service.List(r.Context(), req)
	if err != nil {
		return fmt.Errorf("on List: %w", err)
	}

	_ = json.NewEncoder(w).Encode(deps)

	return nil
}

// Create godoc
// @Summary      Create
// @Description  creates entities
// @Tags         v1
// @Accept       json
// @Produce      json
// @Param        input   body      models.Create  true  "body"
// @Success      201  {int}  0
// @Failure      400  {int}  0
// @Failure      404  {int}  0
// @Failure      500  {int}  0
// @Router       /v1/create [post]
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) error {
	var req models.Create
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return errors.New(err.Error(), errors.ErrBadRequest)
	}

	if err := h.Service.Create(r.Context(), req); err != nil {
		return fmt.Errorf("on Create: %w", err)
	}

	w.WriteHeader(http.StatusCreated)

	return nil
}

// Delete godoc
// @Summary      Delete
// @Description  deletes entities
// @Tags         v1
// @Accept       json
// @Produce      json
// @Param        input   body      models.Delete  true  "body"
// @Success      200  {int}  0
// @Failure      400  {int}  0
// @Failure      404  {int}  0
// @Failure      500  {int}  0
// @Router       /v1/delete [delete]
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) error {
	var req models.Delete
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return errors.New(err.Error(), errors.ErrBadRequest)
	}

	if err := h.Service.Delete(r.Context(), req); err != nil {
		return fmt.Errorf("on Delete: %w", err)
	}

	return nil
}

// Update godoc
// @Summary      Update
// @Description  updates entities
// @Tags         v1
// @Accept       json
// @Produce      json
// @Param        input   body      models.Update  true  "body"
// @Success      200  {int}  0
// @Failure      400  {int}  0
// @Failure      404  {int}  0
// @Failure      500  {int}  0
// @Router       /v1/update [put]
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) error {
	var req models.Update
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return errors.New(err.Error(), errors.ErrBadRequest)
	}

	if err := h.Service.Update(r.Context(), req); err != nil {
		return fmt.Errorf("on Update: %w", err)
	}

	return nil
}
