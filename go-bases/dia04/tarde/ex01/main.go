package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	defer func() {
		fmt.Println("execução concluída")
	}()

	readFile("customers.txt")
}

func readFile(filePath string) (content []byte, err error) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Recovered from panic:", err)
		}
	}()

	if filePath == "" {
		err = errors.New("file path cannot be empty")
		return
	}
	content, err = os.ReadFile(filePath)
	if err != nil {
		err = errors.New("The indicated file was not found or is damaged")
		panic(err)
	}
	return
}
