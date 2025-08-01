package handler

import (
	"app/internal"
	"app/internal/repository"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProductsDefault_Get(t *testing.T) {
	rp := repository.ProductsRepositoryMock{}
	h := NewProductsDefault(&rp)

	t.Run("Success - Search All Products", func(t *testing.T) {
		// given
		expectedCode := http.StatusOK
		expectedMessage := "success"
		products := map[int]internal.Product{
			1: {Id: 1, ProductAttributes: internal.ProductAttributes{Description: "Product 1", Price: 100, SellerId: 1}},
			2: {Id: 2, ProductAttributes: internal.ProductAttributes{Description: "Product 2", Price: 200, SellerId: 2}},
		}
		query := internal.ProductQuery{
			Id: 0,
		}

		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		res := httptest.NewRecorder()

		//when
		rp.On("SearchProducts", query).Return(products, nil)
		h.Get().ServeHTTP(res, req)

		var body map[string]any
		err := json.Unmarshal(res.Body.Bytes(), &body)

		// then
		require.NoError(t, err)
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedMessage, body["message"])
		require.Equal(t, len(products), len(body["data"].(map[string]any)))
		rp.AssertExpectations(t)
	})

	t.Run("Success - Search Product by Id", func(t *testing.T) {
		// given
		expectedCode := http.StatusOK
		expectedMessage := "success"
		products := map[int]internal.Product{
			1: {Id: 1, ProductAttributes: internal.ProductAttributes{Description: "Product 1", Price: 100, SellerId: 1}},
			2: {Id: 2, ProductAttributes: internal.ProductAttributes{Description: "Product 2", Price: 200, SellerId: 2}},
		}
		query := internal.ProductQuery{
			Id: 1,
		}

		req := httptest.NewRequest(http.MethodGet, "/products?id=1", nil)
		res := httptest.NewRecorder()

		//when
		rp.On("SearchProducts", query).Return(map[int]internal.Product{1: products[1]}, nil)
		h.Get().ServeHTTP(res, req)

		var body map[string]any
		err := json.Unmarshal(res.Body.Bytes(), &body)

		// then
		require.NoError(t, err)
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedMessage, body["message"])
		require.Equal(t, 1, len(body["data"].(map[string]any)))
		rp.AssertExpectations(t)
	})

	t.Run("Error - Invalid Id", func(t *testing.T) {
		// given
		expectedCode := http.StatusBadRequest
		expectedMessage := "invalid id"

		req := httptest.NewRequest(http.MethodGet, "/products?id=invalid", nil)
		res := httptest.NewRecorder()

		//when
		h.Get().ServeHTTP(res, req)

		var body map[string]any
		err := json.Unmarshal(res.Body.Bytes(), &body)

		// then
		require.NoError(t, err)
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedMessage, body["message"])
		rp.AssertExpectations(t)
	})

	t.Run("Error - Internal Error", func(t *testing.T) {
		// given
		expectedCode := http.StatusInternalServerError
		expectedMessage := "internal error"
		query := internal.ProductQuery{
			Id: 2,
		}
		products := map[int]internal.Product{}

		req := httptest.NewRequest(http.MethodGet, "/products?id=2", nil)
		res := httptest.NewRecorder()

		//when
		rp.On("SearchProducts", query).Return(products, errors.New("internal error"))
		h.Get().ServeHTTP(res, req)

		var body map[string]any
		err := json.Unmarshal(res.Body.Bytes(), &body)

		// then
		require.NoError(t, err)
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedMessage, body["message"])
		rp.AssertExpectations(t)
	})
}
