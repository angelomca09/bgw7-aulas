package main

import "errors"

func main() {
	salary := 1000

	var err *myError

	if salary < 10000 {
		err = &myError{}
	}

	if errors.Is(err, &myError{}) {
		panic(err)
	}
}

type myError struct{}

func (e *myError) Error() string {
	return "Error: salary is less than 10000"
}
