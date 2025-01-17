package belajar_golang_database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
