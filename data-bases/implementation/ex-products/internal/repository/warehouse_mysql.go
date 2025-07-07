package repository

import (
	"app/internal"
	"database/sql"
)

type mysqlRepositoryWarehouse struct {
	db *sql.DB
}

func NewRepositoryWarehouseMysql(db *sql.DB) internal.RepositoryWarehouse {
	return &mysqlRepositoryWarehouse{
		db: db,
	}
}

// FindById implements internal.RepositoryWarehouse.
func (m *mysqlRepositoryWarehouse) FindById(id int) (p internal.Warehouse, err error) {
	row := m.db.QueryRow(GetOneWarehouseQuery, id)
	if err := row.Err(); err != nil {
		return p, err
	}

	if err := row.Scan(&p.Id, &p.Name, &p.Address, &p.Telephone, &p.Capacity); err != nil {
		return p, err
	}

	return p, nil
}

// Save implements internal.RepositoryWarehouse.
func (m *mysqlRepositoryWarehouse) Save(p *internal.Warehouse) (err error) {
	result, err := m.db.Exec(StoreWarehouseQuery, p.Name, p.Address, p.Telephone, p.Capacity)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	p.Id = int(id)

	return nil
}
