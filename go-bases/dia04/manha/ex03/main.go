package main

import "errors"

func main() {
	salary := 1000

	var err error

	if salary < 10000 {
		err = errors.New("Error: salary is less than 10000")
	}

	if errors.Is(err, err) {
		panic(err)
	}
}
