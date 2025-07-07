package repository_test

import (
	"app/internal/repository"
	"app/tests/utils"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductsFindAll(t *testing.T) {
	utils.RegisterDatabase()
	utils.InitDatabase()
	t.Run("success get all", func(t *testing.T) {
		// arrange
		db := utils.GetDB()
		defer db.Close()

		rp := repository.NewRepositoryProductMysql(db)

		// act
		list, err := rp.FindAll()
		jsonList, _ := json.Marshal(list)

		expectedList := `[{"id":1,"name":"product 1","quantity":10,"code_value":"code 1","is_published":true,"expiration":"2023-01-01T00:00:00Z","price":100,"warehouse_id":1},{"id":2,"name":"product 2","quantity":20,"code_value":"code 2","is_published":true,"expiration":"2023-01-01T00:00:00Z","price":200,"warehouse_id":2},{"id":3,"name":"product 3","quantity":30,"code_value":"code 3","is_published":true,"expiration":"2023-01-01T00:00:00Z","price":300,"warehouse_id":3},{"id":4,"name":"product 4","quantity":40,"code_value":"code 4","is_published":true,"expiration":"2023-01-01T00:00:00Z","price":400,"warehouse_id":4},{"id":5,"name":"product 5","quantity":50,"code_value":"code 5","is_published":true,"expiration":"2023-01-01T00:00:00Z","price":500,"warehouse_id":5}]`

		// assert
		assert.Equal(t, expectedList, string(jsonList))
		assert.Nil(t, err)
	})

}
