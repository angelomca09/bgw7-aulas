package handler_test

import (
	"estrutura-api/internal"
	"estrutura-api/internal/handler"
	"estrutura-api/internal/repository"
	"estrutura-api/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

/* func TestMain(m *testing.M) {

} */

func TestGetAll(t *testing.T) {
	t.Run("GET:/products -> expect success", func(t *testing.T) {
		prs := map[int]*internal.ProductAttributes{
			1: {Name: "Mouse", Quantity: 2, CodeValue: "a", IsPublished: true, Expiration: "10/06/2025", Price: 25.90},
			2: {Name: "Teclado", Quantity: 3, CodeValue: "b", IsPublished: true, Expiration: "10/06/2025", Price: 120.90},
		}
		db := repository.NewProductJSONRepositoryWithMap(prs)
		sv := service.NewProductService(db)
		hd := handler.NewProductHandler(sv)

		req := httptest.NewRequest("GET", "/products", nil)
		res := httptest.NewRecorder()
		hd.GetAll()(res, req)

		expectedCode := http.StatusOK
		expectedBody := `[
        {"id": 1,  "name": "Mouse",  "quantity": 2,  "code_value": "a",  "is_published": true,  "expiration": "10/06/2025",  "price": 25.90},
        {"id": 2,  "name": "Teclado",  "quantity": 3,  "code_value": "b",  "is_published": true,  "expiration": "10/06/2025",  "price": 120.90}
        ]`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())

	})
}

func TestGetById(t *testing.T) {
	t.Run("GET:/products/{id} -> expect success", func(t *testing.T) {
		prs := map[int]*internal.ProductAttributes{
			1: {Name: "Mouse", Quantity: 2, CodeValue: "a", IsPublished: true, Expiration: "10/06/2025", Price: 25.90},
			2: {Name: "Teclado", Quantity: 3, CodeValue: "b", IsPublished: true, Expiration: "10/06/2025", Price: 120.90},
		}
		db := repository.NewProductJSONRepositoryWithMap(prs)
		sv := service.NewProductService(db)
		hd := handler.NewProductHandler(sv)

		req := httptest.NewRequest("GET", "/products/id:1", nil)
		res := httptest.NewRecorder()
		hd.GetById()(res, req)

		//expectedCode := http.StatusOK
		expectedBody := `{"id": 2,  "name": "Mouse",  "quantity": 2,  "code_value": "a",  "is_published": true,  "expiration": "10/06/2025",  "price": 25.90}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		//require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})
}
