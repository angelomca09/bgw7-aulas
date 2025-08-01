package repository

import (
	"app/internal"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProductsMap_NewProductsMap(t *testing.T) {
	t.Run("Create a new ProductsMap without parameters", func(t *testing.T) {
		productsMap := NewProductsMap(nil)
		require.NotNil(t, productsMap)
	})

	t.Run("Create a new ProductsMap with parameters", func(t *testing.T) {
		db := map[int]internal.Product{
			1: {
				Id: 1,
				ProductAttributes: internal.ProductAttributes{
					Description: "Product 1",
					Price:       100,
					SellerId:    1,
				},
			},
			2: {
				Id: 2,
				ProductAttributes: internal.ProductAttributes{
					Description: "Product 2",
					Price:       200,
					SellerId:    2,
				},
			},
		}
		productsMap := NewProductsMap(db)
		require.NotNil(t, productsMap)
	})
}

func TestProductsMap_SearchProducts(t *testing.T) {
	db := map[int]internal.Product{
		1: {
			Id: 1,
			ProductAttributes: internal.ProductAttributes{
				Description: "Product 1",
				Price:       100,
				SellerId:    1,
			},
		},
		2: {
			Id: 2,
			ProductAttributes: internal.ProductAttributes{
				Description: "Product 2",
				Price:       200,
				SellerId:    2,
			},
		},
	}
	t.Run("Search all products", func(t *testing.T) {
		productsMap := NewProductsMap(db)
		products, err := productsMap.SearchProducts(internal.ProductQuery{})
		require.NoError(t, err)
		require.Equal(t, 2, len(products))
		require.Equal(t, 1, products[1].Id)
		require.Equal(t, "Product 1", products[1].Description)
		require.EqualValues(t, 100, products[1].Price)
		require.Equal(t, 1, products[1].SellerId)
	})

	t.Run("Search product with id", func(t *testing.T) {
		productsMap := NewProductsMap(db)
		products, err := productsMap.SearchProducts(internal.ProductQuery{Id: 1})
		require.NoError(t, err)
		require.Equal(t, 1, len(products))
		require.Equal(t, 1, products[1].Id)
		require.Equal(t, "Product 1", products[1].Description)
		require.EqualValues(t, 100, products[1].Price)
		require.Equal(t, 1, products[1].SellerId)
	})
}
