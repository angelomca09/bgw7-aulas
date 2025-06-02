package main

import (
	"errors"
	"fmt"
)

func main() {
	salary := 1000

	var err error

	if salary < 15000 {
		err = fmt.Errorf("Error: the minimum taxable amount is 150,000 and the salary entered is: %d", salary)
	}

	if errors.Is(err, fmt.Errorf("Error: the minimum taxable amount is 150,000 and the salary entered is: %d", salary)) {
		panic(err)
	}
}
