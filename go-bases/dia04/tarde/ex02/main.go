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
	// This may not be necessary
	changeDir("./go-bases/dia04/tarde/ex02")
	// Read using os.ReadFile
	readFile("customers.txt")
	// Read using os.Open
	openAndReadFile("customers.txt")
}

func changeDir(dir string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:", err)
		}
	}()

	err := os.Chdir(dir)
	if err != nil {
		panic(errors.New("Failed to change directory: " + err.Error()))
	}
	fmt.Println("Changed directory to:", dir)
}

func readFile(filePath string) (err error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:", err)
		}
	}()

	if filePath == "" {
		err = errors.New("file path cannot be empty")
		return
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		err = errors.New("The indicated file was not found or is damaged")
		panic(err)
	}

	fmt.Println("File read successfully:\n", (string)(content))
	return
}

func openAndReadFile(filePath string) (err error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:", err)
		}
	}()

	if filePath == "" {
		err = errors.New("file path cannot be empty")
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		err = errors.New("The indicated file was not found or is damaged")
		panic(err)
	}
	defer file.Close()

	stats, err := file.Stat()
	if err != nil {
		err = errors.New("Failed to get file stats")
		panic(err)
	}

	content := make([]byte, stats.Size())
	_, err = file.Read(content)

	fmt.Println("File read successfully:\n", (string)(content))

	return
}
