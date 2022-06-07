package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Arguments : ")
	fmt.Println(args)

	//fmt.Println(args[1])
	//fmt.Println(args[2])
	//fmt.Println(args[3])

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname : ", name)
	} else {
		fmt.Println("Error", err.Error())
	}

	//username := os.Getenv("APP_USERNAME")
	//password := os.Getenv("APP_PASSWORD")
	//
	//fmt.Println(username)
	//fmt.Println(password)

	username := os.Getenv("TEST-DATABASE")
	password := os.Getenv("TEST-PASSWORD")

	fmt.Println(username)
	fmt.Println(password)
}
