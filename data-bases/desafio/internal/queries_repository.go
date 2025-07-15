package internal

// QueriesRepository is the interface that wraps the basic methods that a queries repository should implement.
type QueriesRepository interface {
	// UpdateInvoiceTotals updates the total of all invoices
	UpdateInvoiceTotals() (err error)
	// GetTotalByCondition returns the total by condition
	GetTotalByCondition() (t []TotalByCondition, err error)
	// GetTop5SoldProducts returns the top 5 sold products
	GetTop5SoldProducts() (t []Top5SoldProducts, err error)
	// GetTop5ActiveCustomersByTotalSpent returns the top 5 active customers by total spent
	GetTop5ActiveCustomersByTotalSpent() (t []Top5ActiveCustomersByTotalSpent, err error)
}
