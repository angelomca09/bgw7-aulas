package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Initialize() (err error) {
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/products_db")
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	log.Println("Connected to MySQL database")

	return nil
}

func GetDB() *sql.DB {
	return db
}
