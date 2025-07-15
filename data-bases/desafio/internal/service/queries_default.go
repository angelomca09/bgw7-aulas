package service

import "app/internal"

// NewQueriesDefault creates new default service for queries.
func NewQueriesDefault(rp internal.QueriesRepository) *QueriesDefault {
	return &QueriesDefault{rp}
}

// QueriesDefault is the default service implementation for queries.
type QueriesDefault struct {
	// rp is the repository for queries.
	rp internal.QueriesRepository
}

// UpdateInvoiceTotals updates the invoice totals.
func (s *QueriesDefault) UpdateInvoiceTotals() (err error) {
	return s.rp.UpdateInvoiceTotals()
}

// GetTotalByCondition returns the total by condition.
func (s *QueriesDefault) GetTotalByCondition() (t []internal.TotalByCondition, err error) {
	return s.rp.GetTotalByCondition()
}

// GetTop5SoldProducts returns the top 5 sold products.
func (s *QueriesDefault) GetTop5SoldProducts() (t []internal.Top5SoldProducts, err error) {
	return s.rp.GetTop5SoldProducts()
}

// GetTop5ActiveCustomersByTotalSpent returns the top 5 active customers by total spent.
func (s *QueriesDefault) GetTop5ActiveCustomersByTotalSpent() (t []internal.Top5ActiveCustomersByTotalSpent, err error) {
	return s.rp.GetTop5ActiveCustomersByTotalSpent()
}
