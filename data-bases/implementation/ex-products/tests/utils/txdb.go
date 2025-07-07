package utils

import (
	"database/sql"

	"github.com/DATA-DOG/go-txdb"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func RegisterDatabase() {

	cfg := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "products_test_db",
	}
	txdb.Register("txdb", "mysql", cfg.FormatDSN())
}

func InitDatabase() error {
	conn, err := sql.Open("txdb", "identifier")

	db = conn

	return err
}

func GetDB() *sql.DB {
	return db
}
