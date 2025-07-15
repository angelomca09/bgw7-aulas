package main

import (
	"app/internal/application"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// env
	// ...

	// app
	// - config
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, using system environment variables")
	}

	cfg := &application.ConfigApplicationDefault{
		Db: &mysql.Config{
			User:   getEnv("DB_USER", ""),
			Passwd: getEnv("DB_PASSWORD", ""),
			Net:    "tcp",
			Addr:   fmt.Sprintf("%s:%s", getEnv("DB_HOST", ""), getEnv("DB_PORT", "")),
			DBName: getEnv("DB_NAME", ""),
		},
		Addr: fmt.Sprintf("%s:%s", getEnv("HOST", ""), getEnv("PORT", "8080")),
	}
	app := application.NewApplicationDefault(cfg)
	// - set up
	err := app.SetUp()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Application all set up and ready to go!")
	// - run
	fmt.Println("Running application")
	err = app.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
