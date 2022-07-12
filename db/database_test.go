package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

//TODO: TESTING CONNECTION DATABASE

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_cms_golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//TODO: DEFER FUNCTION KE 2 UNTUK GENERATE SQL

	//defer func(db *sql.DB) {
	//	err := db.Close()
	//	if err != nil {
	//
	//	}
	//}(db)
}

// TODO: DATABASE POOLING
