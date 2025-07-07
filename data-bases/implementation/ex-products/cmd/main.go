package main

import (
	"app/internal/application"
	"app/internal/database"
	"fmt"
)

func main() {
	// env
	err := database.Initialize()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer database.GetDB().Close()

	// app
	// - config
	app := application.NewApplicationDatabase("")
	// - tear down
	defer app.TearDown()
	// - set up
	if err := app.SetUp(); err != nil {
		fmt.Println(err)
		return
	}
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
