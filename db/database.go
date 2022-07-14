package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// TODO: PARSE TIME UNTUK MELAKUKAN PEMBACAAN OTOMATIS DI DRIVER MYSQL GOLANG ( AUTOMATIC PARSING )
func GetConnection() (*sql.DB, *sql.DB) {
	db1, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_clife_epay?parseTime=true")
	db2, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_cms_golang?parseTime=true")

	if err != nil {
		panic(err)
	}

	db1.SetMaxIdleConns(10)
	db1.SetMaxOpenConns(100)
	db1.SetConnMaxIdleTime(5 * time.Minute)
	db1.SetConnMaxLifetime(60 * time.Minute)

	db2.SetMaxIdleConns(10)
	db2.SetMaxOpenConns(100)
	db2.SetConnMaxIdleTime(5 * time.Minute)
	db2.SetConnMaxLifetime(60 * time.Minute)

	return db1, db2
}

func ConnectionGOCMS() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_golang?parseTime=true")

	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
