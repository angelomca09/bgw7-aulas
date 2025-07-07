package repository

import (
	"app/internal"
	"database/sql"
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
	if err := row.Scan(&p.Id, &p.Name, &p.Quantity, &p.CodeValue, &p.IsPublished, &expirationStr, &p.Price); err != nil {

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
	result, err := m.db.Exec(StoreProductQuery, p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration, p.Price)
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

// UpdateOrSave implements internal.RepositoryProduct.
func (m *mysqlRepository) UpdateOrSave(p *internal.Product) (err error) {
	panic("unimplemented")
}

// Update implements internal.RepositoryProduct.
func (m *mysqlRepository) Update(p *internal.Product) (err error) {
	_, err = m.db.Exec(UpdateProductQuery, p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration, p.Price, p.Id)
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
