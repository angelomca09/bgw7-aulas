package internal

type QueriesService interface {
	// UpdateInvoiceTotals updates the total of all invoices
	UpdateInvoiceTotals() (err error)
	// TotalByCondition returns the total amount of invoices by condition.
	GetTotalByCondition() (t []TotalByCondition, err error)
	// Top5SoldProducts returns the top 5 sold products.
	GetTop5SoldProducts() (t []Top5SoldProducts, err error)
	// Top5ActiveCustomersByTotalSpent returns the top 5 active customers by total spent
	GetTop5ActiveCustomersByTotalSpent() (t []Top5ActiveCustomersByTotalSpent, err error)
}
