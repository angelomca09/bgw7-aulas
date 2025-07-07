package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"app/tests/utils"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWareHouseFindAll(t *testing.T) {
	utils.RegisterDatabase()
	utils.InitDatabase()
	t.Run("success get all", func(t *testing.T) {
		// arrange
		db := utils.GetDB()
		defer db.Close()

		rp := repository.NewRepositoryWarehouseMysql(db)

		// act
		list, err := rp.FindAll()
		jsonList, _ := json.Marshal(list)

		expectedList := `[{"id":1,"name":"Main Warehouse","address":"221 Baker Street","telephone":"4555666","capacity":100},{"id":2,"name":"Secondary Warehouse","address":"334 Baker Street","telephone":"7888999","capacity":150},{"id":3,"name":"Third Warehouse","address":"334 Baker Street","telephone":"7888999","capacity":150},{"id":4,"name":"Fourth Warehouse","address":"334 Baker Street","telephone":"7888999","capacity":150},{"id":5,"name":"Fifth Warehouse","address":"334 Baker Street","telephone":"7888999","capacity":150}]`

		// assert
		assert.Equal(t, expectedList, string(jsonList))
		assert.Nil(t, err)
	})

	t.Run("success get by id", func(t *testing.T) {
		// arrange
		db := utils.GetDB()
		defer db.Close()

		rp := repository.NewRepositoryWarehouseMysql(db)

		// act
		list, err := rp.FindById(1)
		jsonList, _ := json.Marshal(list)

		expectedList := `{"id":1,"name":"Main Warehouse","address":"221 Baker Street","telephone":"4555666","capacity":100}`

		// assert
		assert.Equal(t, expectedList, string(jsonList))
		assert.Nil(t, err)
	})

	t.Run("success create", func(t *testing.T) {
		// arrange
		db := utils.GetDB()
		defer db.Close()

		rp := repository.NewRepositoryWarehouseMysql(db)

		// act
		err := rp.Save(&internal.Warehouse{
			Id: 6,
			WarehouseAttributes: internal.WarehouseAttributes{
				Name:      "Warehouse 6",
				Address:   "Address 6",
				Telephone: "Telephone 6",
				Capacity:  600,
			},
		})

		// assert
		assert.Nil(t, err)
	})
}
