package repository

import (
	"app/internal"
	"database/sql"
)

func NewQueriesRepositoryMySQL(db *sql.DB) *QueriesRepositoryMySQL {
	return &QueriesRepositoryMySQL{db}
}

type QueriesRepositoryMySQL struct {
	db *sql.DB
}

func (r *QueriesRepositoryMySQL) UpdateInvoiceTotals() (err error) {
	_, err = r.db.Exec(UpdateInvoiceTotals)
	return // Will return error if any is returned by db.Exec
}

func (r *QueriesRepositoryMySQL) GetTotalByCondition() (t []internal.TotalByCondition, err error) {
	t = make([]internal.TotalByCondition, 0)
	rows, err := r.db.Query(GetTotalByCondition)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var total internal.TotalByCondition
		err := rows.Scan(&total.Condition, &total.Total)
		if err != nil {
			return nil, err
		}
		t = append(t, total)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return
}

func (r *QueriesRepositoryMySQL) GetTop5SoldProducts() (t []internal.Top5SoldProducts, err error) {
	t = make([]internal.Top5SoldProducts, 0)
	rows, err := r.db.Query(GetTop5SoldProducts)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var top internal.Top5SoldProducts
		err := rows.Scan(&top.Description, &top.Total)
		if err != nil {
			return nil, err
		}
		t = append(t, top)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return
}

func (r *QueriesRepositoryMySQL) GetTop5ActiveCustomersByTotalSpent() (t []internal.Top5ActiveCustomersByTotalSpent, err error) {
	t = make([]internal.Top5ActiveCustomersByTotalSpent, 0)
	rows, err := r.db.Query(GetTop5ActiveCustomersByTotalSpent)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var top internal.Top5ActiveCustomersByTotalSpent
		err := rows.Scan(&top.FirstName, &top.LastName, &top.Amount)
		if err != nil {
			return nil, err
		}
		t = append(t, top)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return
}
