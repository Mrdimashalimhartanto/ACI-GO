package main

import (
	"context"
	"fmt"
	"testing"
)

func TestSqlInjectionTest(t *testing.T) {
	db := DemodbConnection()
	defer db.Close()

	ctx := context.Background()

	// TODO: USERNAME PASSWORD ADMIN DARI DATABASE
	//**
	// tanda ; dan # ( crash ) tanda pada sql injection untuk melakukan sebuah
	// method injection dengan bypass sebuah password
	//*/
	username := "admin'; #"
	password := "admin"

	script := "SELECT username FROM user WHERE username = '" + username +
		"' AND password = '" + password + "' LIMIT  1"
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}
