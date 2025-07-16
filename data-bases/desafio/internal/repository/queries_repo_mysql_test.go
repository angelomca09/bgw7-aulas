package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"app/test/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueriesRepoMysql_UpdateInvoiceTotals(t *testing.T) {
	utils.RegisterDatabase()
	utils.InitDatabase()
	t.Run("should Update Invoice Totals successfully", func(t *testing.T) {
		// Arrange
		db := utils.GetDB()
		defer db.Close()

		queriesRepo := repository.NewQueriesRepositoryMySQL(db)

		// Act
		err := queriesRepo.UpdateInvoiceTotals()

		// Assert
		assert.NoError(t, err)
	})
}

func TestQueriesRepoMysql_GetTotalCondition(t *testing.T) {
	utils.RegisterDatabase()
	utils.InitDatabase()
	t.Run("should return Totals By Condition", func(t *testing.T) {
		// Arrange
		db := utils.GetDB()
		defer db.Close()

		queriesRepo := repository.NewQueriesRepositoryMySQL(db)
		expectedTotals := []internal.TotalByCondition{
			{Condition: 1, Total: 8870},
			{Condition: 0, Total: 10470},
		}

		// Act
		totals, err := queriesRepo.GetTotalByCondition()

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, totals)
		assert.NotEmpty(t, totals)
		assert.Equal(t, expectedTotals, totals)
	})
}

func TestQueriesRepoMysql_Top5Products(t *testing.T) {
	utils.RegisterDatabase()
	utils.InitDatabase()
	t.Run("should return Top 5 Products", func(t *testing.T) {
		// Arrange
		db := utils.GetDB()
		defer db.Close()

		queriesRepo := repository.NewQueriesRepositoryMySQL(db)
		expectedProducts := []internal.Top5SoldProducts{
			{Description: "Product 4", Total: 200},
			{Description: "Product 3", Total: 113},
			{Description: "Product 6", Total: 70},
			{Description: "Product 1", Total: 50},
			{Description: "Product 2", Total: 50},
		}

		// Act
		products, err := queriesRepo.GetTop5SoldProducts()

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, products)
		assert.NotEmpty(t, products)
		assert.Equal(t, expectedProducts, products)
	})
}

func TestQueriesRepoMysql_Top5Customers(t *testing.T) {
	utils.RegisterDatabase()
	utils.InitDatabase()
	t.Run("should return Top 5 Customers", func(t *testing.T) {
		// Arrange
		db := utils.GetDB()
		defer db.Close()

		queriesRepo := repository.NewQueriesRepositoryMySQL(db)
		expectedCustomers := []internal.Top5ActiveCustomersByTotalSpent{
			{FirstName: "Eve", LastName: "Johnson", Amount: 3910},
			{FirstName: "David", LastName: "Jones", Amount: 3530},
			{FirstName: "Frank", LastName: "Davis", Amount: 2680},
			{FirstName: "Bob", LastName: "Smith", Amount: 2580},
			{FirstName: "Charlie", LastName: "Brown", Amount: 2080},
		}

		// Act
		customers, err := queriesRepo.GetTop5ActiveCustomersByTotalSpent()

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, customers)
		assert.NotEmpty(t, customers)
		assert.Equal(t, expectedCustomers, customers)
	})
}
