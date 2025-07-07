package handler

import (
	"app/internal"
	"app/platform/web/request"
	"app/platform/web/response"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// NewHandlerWarehouse creates a new handler for warehosue.
func NewHandlerWarehouse(rp internal.RepositoryWarehouse) (h *HandlerWarehouse) {
	h = &HandlerWarehouse{
		rp: rp,
	}
	return
}

// HandlerWarehouse is a handler for warehouse.
type HandlerWarehouse struct {
	// rp is the repository for warehouse.
	rp internal.RepositoryWarehouse
}

// GetById gets a warehouse by id.
func (h *HandlerWarehouse) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - path parameter: id
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "invalid id")
			return
		}

		// process
		// - find warehouse by id
		wa, err := h.rp.FindById(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrRepositoryWarehouseNotFound):
				response.JSON(w, http.StatusNotFound, "warehouse not found")
			default:
				fmt.Println("error: ", err.Error())
				response.JSON(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    wa,
		})
	}
}

// Create creates a warehouse.
func (h *HandlerWarehouse) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - body
		var body internal.Warehouse
		err := request.JSON(r, &body)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "invalid body")
			return
		}

		err = h.rp.Save(&body)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, "internal server error")
			return
		}

		// response
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "success",
			"data":    body,
		})
	}
}
