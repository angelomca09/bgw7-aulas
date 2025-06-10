package handler

import (
	"encoding/json"
	"estrutura-api/internal"
	"estrutura-api/platform/web/request"
	"estrutura-api/platform/web/response"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	service internal.ProductService
}

func NewProductHandler(s internal.ProductService) *ProductHandler {
	hd := &ProductHandler{}

	hd.service = s

	return hd
}

func (hd *ProductHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := hd.service.GetAll()
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		response.JSON(w, http.StatusOK, products)
	}
}

func (hd *ProductHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		product, err := hd.service.GetByID(id)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		response.JSON(w, http.StatusOK, product)
	}
}

func (hd *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody := &internal.ProductAttributes{}
		err := request.JSON(r, reqBody)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid body")
			return
		}

		pr, err := hd.service.Create(reqBody)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "Error Creating the Product: "+err.Error())
			return
		}

		response.JSON(w, http.StatusOK, pr)
	}
}

func (hd *ProductHandler) UpdateOrCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid id")
			return
		}

		reqBody := &internal.ProductAttributes{}

		err = request.JSON(r, reqBody)
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		pr := &internal.Product{
			ID: id, ProductAttributes: *reqBody,
		}

		pr, err = hd.service.UpdateOrCreate(pr)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		response.JSON(w, http.StatusOK, pr)

	}
}

func (hd *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid id")
			return
		}

		pr, err := hd.service.GetByID(id)
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		reqBody := &pr.ProductAttributes

		err = request.JSON(r, reqBody)
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		pr = &internal.Product{ID: id, ProductAttributes: *reqBody}

		pr, err = hd.service.Update(id, pr)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		response.JSON(w, http.StatusOK, pr)
	}
}

func (hd *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		if _, err = hd.service.GetByID(id); err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		err = hd.service.Delete(id)
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		response.Text(w, http.StatusOK, "Product removed")
	}
}

func (hd *ProductHandler) GetTotalPriceByList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		listString := r.URL.Query().Get("list")
		var list []int

		if listString != "" {
			err := json.Unmarshal([]byte(listString), &list)
			if err != nil {
				response.Error(w, http.StatusBadGateway, err.Error())
				return
			}
		}

		prodList, totalPrice, err := hd.service.GetTotalPriceByList(list)
		if err != nil {
			response.Error(w, http.StatusBadGateway, err.Error())
		}

		body := struct {
			ProdList   []*internal.Product `json:"prodList"`
			TotalPrice float64             `json:"total_price"`
		}{
			ProdList:   prodList,
			TotalPrice: totalPrice,
		}

		response.JSON(w, http.StatusOK, body)
	}
}
