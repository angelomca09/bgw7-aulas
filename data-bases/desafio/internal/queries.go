package internal

type TotalByCondition struct {
	Condition int     `json:"condition"`
	Total     float64 `json:"total"`
}

type Top5SoldProducts struct {
	Description string `json:"description"`
	Total       int    `json:"total"`
}

type Top5ActiveCustomersByTotalSpent struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Amount    float64 `json:"amount"`
}
