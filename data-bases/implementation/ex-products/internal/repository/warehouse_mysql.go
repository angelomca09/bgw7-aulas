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

// FindAll implements internal.RepositoryWarehouse.
func (m *mysqlRepositoryWarehouse) FindAll() (w_list []internal.Warehouse, err error) {
	rows, err := m.db.Query(GetAllWarehousesQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item internal.Warehouse
		if err := rows.Scan(&item.Id, &item.Name, &item.Address, &item.Telephone, &item.Capacity); err != nil {
			if err == sql.ErrNoRows {
				return nil, internal.ErrRepositoryWarehouseNotFound
			}

			return nil, err
		}
		w_list = append(w_list, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return w_list, nil
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
