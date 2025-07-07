package repository

import (
	"app/internal"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type mysqlRepository struct {
	db *sql.DB
}

func NewRepositoryProductMysql(db *sql.DB) internal.RepositoryProduct {
	return &mysqlRepository{
		db: db,
	}
}

// FindById implements internal.RepositoryProduct.
func (m *mysqlRepository) FindById(id int) (p internal.Product, err error) {
	row := m.db.QueryRow(GetOneProductQuery, id)
	if err := row.Err(); err != nil {
		return p, err
	}

	var expirationStr sql.NullString
	if err := row.Scan(&p.Id, &p.Name, &p.Quantity, &p.CodeValue, &p.IsPublished, &expirationStr, &p.Price, &p.WarehouseId); err != nil {

		if err == sql.ErrNoRows {
			return p, internal.ErrRepositoryProductNotFound
		}

		return p, err
	}

	// Parse the expiration date string to time.Time
	if expirationStr.Valid {
		if parsedTime, parseErr := time.Parse("2006-01-02 15:04:05", expirationStr.String); parseErr == nil {
			p.Expiration = parsedTime
		} else if parsedTime, parseErr := time.Parse("2006-01-02", expirationStr.String); parseErr == nil {
			p.Expiration = parsedTime
		}
		// If parsing fails, Expiration remains zero value
	}

	return p, nil

}

// Save implements internal.RepositoryProduct.
func (m *mysqlRepository) Save(p *internal.Product) (err error) {
	result, err := m.db.Exec(StoreProductQuery, p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration, p.Price, p.WarehouseId)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	(*p).Id = int(id)
	return nil
}

// Update implements internal.RepositoryProduct.
func (m *mysqlRepository) Update(p *internal.Product) (err error) {
	_, err = m.db.Exec(UpdateProductQuery, p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration, p.Price, p.WarehouseId, p.Id)
	if err != nil {
		return err
	}

	return nil
}

// Delete implements internal.RepositoryProduct.
func (m *mysqlRepository) Delete(id int) (err error) {
	_, err = m.db.Exec(DeleteProductQuery, id)
	if err != nil {
		return err
	}

	return nil
}

// GetTotalQuantityByWarehouse implements internal.RepositoryProduct.
func (m *mysqlRepository) GetTotalQuantityByWarehouse(id int) (list []internal.QuantityByWarehouse, err error) {

	querry := GetInnerJoinProductWarehouseQuery

	log.Println("id: ", id)
	if id != 0 {
		querry += fmt.Sprintf(" WHERE w.id = %d", id)
	} else {
		querry += " GROUP BY w.id ORDER BY w.id"
	}

	rows, err := m.db.Query(querry)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item internal.QuantityByWarehouse
		if err := rows.Scan(&item.Name, &item.Quantity); err != nil {
			return nil, err
		}
		list = append(list, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}
