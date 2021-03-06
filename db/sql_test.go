package main

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

//TODO: INSERT QUERY FROM MYSQL

func TestExceSql(t *testing.T) {

	db1 := GetConnection()
	defer db1.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('1', 'Dimas')"
	_, err := db1.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success menambahkan data")

}

//TODO: SELECT QUERY FROM MYSQL

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
	}
}

// TODO: TEST QUERY SQL COMPLETE
func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, response, title, message, created_at FROM logs"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id float64
		var response string
		var title string
		var message string
		var createdAt time.Time

		err = rows.Scan(&id, &response, &title, &message, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("==================================")
		fmt.Println(id)
		fmt.Println(title)
		fmt.Println(response)
		fmt.Println(createdAt)
	}
}

/**
TODO: maping tipe data golang
===============================================
TODO: TIPE DATA DATABASE GOLANG

TODO: 1. VARCHAR, CHAR  -- tipe data golang = string,
TODO: 2. INT, BIGINT -- tipe data golang = int32, int64,
TODO: 3. FLOAT, DOUBLE -- tipe data golang = float32, float64,
TODO: 4. BOOLEAN -- tipe data golang = bool,
TODO: 5. DATE, DATETIME, TIME, TIMESTAMP -- tipe data golang = time.Time

*/

func TestDBGolang(t *testing.T) {
	db := DemodbConnection()
	defer db.Close()

	ctx := context.Background()

	//TODO: DATA INI HARUS URUT SAAT MELAKUKAN EXECUTE QUERY DARI DATABASE
	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at  FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	//TODO: defer rows digunakan untuk melakukan stop pada hasil query
	//** TODO: tipe data nullable dalam golang
	//	tipe data golang		tipe data nullable
	// 1. tipe data string ----- database/sql.NullString
	// 2. bool --- database/sql.NullBool
	// 3. float64 --- database/sql.NullFloat64
	//*/

	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		// TODO: DATA INI HARUS SAMA DENGAN DATA EXECUTE
		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("===================================")
		fmt.Println(id)
		fmt.Println(name)
		if email.Valid {
			fmt.Println("Email:", email.String)
		}
		fmt.Println(email)

		if birthDate.Valid {
			fmt.Println("Birth Date:", birthDate.Time)
		}

		fmt.Println(married)
		fmt.Println("Created At", createdAt)
	}
}
