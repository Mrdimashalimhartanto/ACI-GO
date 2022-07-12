package main

import (
	"context"
	"fmt"
	"testing"
)

func TestExceSql(t *testing.T) {

	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('2', 'Dimas')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success menambahkan data")

}
