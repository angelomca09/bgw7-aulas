package utils

import (
	"estrutura-api/internal"
	"fmt"
	"time"
)

func ValidateEmptyProduct(p *internal.ProductAttributes) error {
	if p.Name == "" || p.Quantity == 0 || p.Price == 0 ||
		p.CodeValue == "" || p.Expiration == "" {
		return fmt.Errorf("Invalid Values")
	}
	return nil
}

func ValidateCodeValue(cv string, prs []*internal.Product) error {
	if cv == "" {
		return fmt.Errorf("CodeValue is empty")
	}
	for _, pr := range prs {
		if pr.CodeValue == cv {
			return fmt.Errorf("CodeValue already exists")
		}
	}
	return nil
}

func ValidateExpiration(expiration string) error {
	if expiration == "" {
		return fmt.Errorf("Expiration is empty")
	}
	_, err := time.Parse("02/01/2006", expiration)
	if err != nil {
		return fmt.Errorf("Expiration is an invalid date")
	}
	return nil
}
