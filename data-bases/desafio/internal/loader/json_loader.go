package loader

import (
	"app/internal"
	"encoding/json"
	"fmt"
	"os"
)

// DataLoader is responsible for loading data from JSON files into the database
type DataLoader struct {
	customerRepo internal.RepositoryCustomer
	productRepo  internal.RepositoryProduct
	invoiceRepo  internal.RepositoryInvoice
	saleRepo     internal.RepositorySale
}

// NewDataLoader creates a new DataLoader instance
func NewDataLoader(
	customerRepo internal.RepositoryCustomer,
	productRepo internal.RepositoryProduct,
	invoiceRepo internal.RepositoryInvoice,
	saleRepo internal.RepositorySale,
) *DataLoader {
	return &DataLoader{
		customerRepo: customerRepo,
		productRepo:  productRepo,
		invoiceRepo:  invoiceRepo,
		saleRepo:     saleRepo,
	}
}

// CustomerJSON represents the JSON structure for customer data
type CustomerJSON struct {
	ID        int    `json:"id"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Condition int    `json:"condition"`
}

// ProductJSON represents the JSON structure for product data
type ProductJSON struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// InvoiceJSON represents the JSON structure for invoice data
type InvoiceJSON struct {
	ID         int     `json:"id"`
	Datetime   string  `json:"datetime"`
	CustomerID int     `json:"customer_id"`
	Total      float64 `json:"total"`
}

// SaleJSON represents the JSON structure for sale data
type SaleJSON struct {
	ID        int `json:"id"`
	ProductID int `json:"product_id"`
	InvoiceID int `json:"invoice_id"`
	Quantity  int `json:"quantity"`
}

// LoadCustomersFromJSON reads customers from JSON file and saves them to database
func (dl *DataLoader) LoadCustomersFromJSON(filePath string) error {
	customerList, err := dl.customerRepo.FindAll()
	if err != nil {
		return fmt.Errorf("failed to find all customers: %w", err)
	}
	if len(customerList) > 0 {
		return fmt.Errorf("customers already exist in the database")
	}

	// Read JSON file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read customers JSON file: %w", err)
	}

	// Parse JSON
	var customersJSON []CustomerJSON
	if err := json.Unmarshal(data, &customersJSON); err != nil {
		return fmt.Errorf("failed to parse customers JSON: %w", err)
	}

	// Convert and save to database
	for _, customerJSON := range customersJSON {
		customer := internal.Customer{
			Id: customerJSON.ID,
			CustomerAttributes: internal.CustomerAttributes{
				FirstName: customerJSON.FirstName,
				LastName:  customerJSON.LastName,
				Condition: customerJSON.Condition,
			},
		}

		if err := dl.customerRepo.Save(&customer); err != nil {
			return fmt.Errorf("failed to save customer with ID %d: %w", customer.Id, err)
		}
	}

	return nil
}

// LoadProductsFromJSON reads products from JSON file and saves them to database
func (dl *DataLoader) LoadProductsFromJSON(filePath string) error {
	productsList, err := dl.productRepo.FindAll()
	if err != nil {
		return fmt.Errorf("failed to find all products: %w", err)
	}
	if len(productsList) > 0 {
		return fmt.Errorf("products already exist in the database")
	}

	// Read JSON file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read products JSON file: %w", err)
	}

	// Parse JSON
	var productsJSON []ProductJSON
	if err := json.Unmarshal(data, &productsJSON); err != nil {
		return fmt.Errorf("failed to parse products JSON: %w", err)
	}

	// Convert and save to database
	for _, productJSON := range productsJSON {
		product := internal.Product{
			Id: productJSON.ID,
			ProductAttributes: internal.ProductAttributes{
				Description: productJSON.Description,
				Price:       productJSON.Price,
			},
		}

		if err := dl.productRepo.Save(&product); err != nil {
			return fmt.Errorf("failed to save product with ID %d: %w", product.Id, err)
		}
	}

	return nil
}

// LoadInvoicesFromJSON reads invoices from JSON file and saves them to database
func (dl *DataLoader) LoadInvoicesFromJSON(filePath string) error {
	invoicesList, err := dl.invoiceRepo.FindAll()
	if err != nil {
		return fmt.Errorf("failed to find all invoices: %w", err)
	}
	if len(invoicesList) > 0 {
		return fmt.Errorf("invoices already exist in the database")
	}

	// Read JSON file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read invoices JSON file: %w", err)
	}

	// Parse JSON
	var invoicesJSON []InvoiceJSON
	if err := json.Unmarshal(data, &invoicesJSON); err != nil {
		return fmt.Errorf("failed to parse invoices JSON: %w", err)
	}

	// Convert and save to database
	for _, invoiceJSON := range invoicesJSON {
		invoice := internal.Invoice{
			Id: invoiceJSON.ID,
			InvoiceAttributes: internal.InvoiceAttributes{
				Datetime:   invoiceJSON.Datetime,
				Total:      invoiceJSON.Total,
				CustomerId: invoiceJSON.CustomerID,
			},
		}

		if err := dl.invoiceRepo.Save(&invoice); err != nil {
			return fmt.Errorf("failed to save invoice with ID %d: %w", invoice.Id, err)
		}
	}

	return nil
}

// LoadSalesFromJSON reads sales from JSON file and saves them to database
func (dl *DataLoader) LoadSalesFromJSON(filePath string) error {
	salesList, err := dl.saleRepo.FindAll()
	if err != nil {
		return fmt.Errorf("failed to find all sales: %w", err)
	}
	if len(salesList) > 0 {
		return fmt.Errorf("sales already exist in the database")
	}

	// Read JSON file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read sales JSON file: %w", err)
	}

	// Parse JSON
	var salesJSON []SaleJSON
	if err := json.Unmarshal(data, &salesJSON); err != nil {
		return fmt.Errorf("failed to parse sales JSON: %w", err)
	}

	// Convert and save to database
	for _, saleJSON := range salesJSON {
		sale := internal.Sale{
			Id: saleJSON.ID,
			SaleAttributes: internal.SaleAttributes{
				Quantity:  saleJSON.Quantity,
				ProductId: saleJSON.ProductID,
				InvoiceId: saleJSON.InvoiceID,
			},
		}

		if err := dl.saleRepo.Save(&sale); err != nil {
			return fmt.Errorf("failed to save sale with ID %d: %w", sale.Id, err)
		}
	}

	return nil
}

// LoadAllDataFromJSON loads all data from JSON files in the specified directory
func (dl *DataLoader) LoadAllDataFromJSON(jsonDir string) error {
	// Load in the correct order to respect foreign key constraints
	// 1. Customers first (no dependencies)
	if err := dl.LoadCustomersFromJSON(jsonDir + "/customers.json"); err != nil {
		return fmt.Errorf("failed to load customers: %w", err)
	}
	fmt.Println("✓ Customers loaded successfully")

	// 2. Products (no dependencies)
	if err := dl.LoadProductsFromJSON(jsonDir + "/products.json"); err != nil {
		return fmt.Errorf("failed to load products: %w", err)
	}
	fmt.Println("✓ Products loaded successfully")

	// 3. Invoices (depends on customers)
	if err := dl.LoadInvoicesFromJSON(jsonDir + "/invoices.json"); err != nil {
		return fmt.Errorf("failed to load invoices: %w", err)
	}
	fmt.Println("✓ Invoices loaded successfully")

	// 4. Sales (depends on products and invoices)
	if err := dl.LoadSalesFromJSON(jsonDir + "/sales.json"); err != nil {
		return fmt.Errorf("failed to load sales: %w", err)
	}
	fmt.Println("✓ Sales loaded successfully")

	fmt.Println("✓ All data loaded successfully!")
	return nil
}
