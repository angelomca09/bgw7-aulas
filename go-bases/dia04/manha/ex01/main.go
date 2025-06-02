package main

import "fmt"

func main() {
	salary := 123456

	if salary < 150000 {
		err := &myError{}
		panic(err)
	} else {
		fmt.Println("Must pay taxes")
	}
}

type myError struct{}

func (e *myError) Error() string {
	return "Error: the salary entered does not reach the taxable minimum"
}
