package handler

import (
	"app/internal"
	"net/http"

	"github.com/bootcamp-go/web/response"
)

func NewQueriesDefault(sv internal.QueriesService) *QueriesDefault {
	return &QueriesDefault{sv: sv}
}

type QueriesDefault struct {
	sv internal.QueriesService
}

func (h *QueriesDefault) UpdateInvoiceTotals() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// process
		err := h.sv.UpdateInvoiceTotals()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error updating invoice totals")
			return
		}

		// response
		response.JSON(w, http.StatusOK, "invoice totals updated successfully")
	}
}

func (h *QueriesDefault) GetTotalByCondition() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// process
		t, err := h.sv.GetTotalByCondition()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error getting total by condition")
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "total by condition found",
			"data":    t,
		})
	}
}

func (h *QueriesDefault) GetTop5SoldProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// process
		t, err := h.sv.GetTop5SoldProducts()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error getting top 5 sold products")
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "top 5 sold products found",
			"data":    t,
		})
	}
}

func (h *QueriesDefault) GetTop5ActiveCustomersByTotalSpent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// process
		t, err := h.sv.GetTop5ActiveCustomersByTotalSpent()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error getting top 5 active customers by total spent")
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "top 5 active customers by total spent found",
			"data":    t,
		})
	}
}
