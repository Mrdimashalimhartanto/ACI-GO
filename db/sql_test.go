package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

//TODO: INSERT QUERY FROM MYSQL

func TestExceSql(t *testing.T) {

	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('1', 'Dimas')"
	_, err := db.ExecContext(ctx, script)
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
func TestQuerySqlComplete(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, title, order_id, no_polis, action, type, request, response, message, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, title, order_id string
		var balance int32
		var rating float64
		var birthDate, createdAt time.Time
		var married bool

		err = rows.Scan(&id, &title, &order_id, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("==================================")
		fmt.Println("Id:", id)
		fmt.Println("Title:", title)
		fmt.Println("Order ID:", order_id)
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		fmt.Println("Birth Date:", birthDate)
		fmt.Println("Married:", married)
		fmt.Println("Created At:", createdAt)
	}
}
