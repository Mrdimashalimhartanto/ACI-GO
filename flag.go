package main

import (
	"flag"
	"fmt"
)

// package flag

func main() {

	var host = flag.String("host", "localhost", "Put your database host")
	var user = flag.String("user", "root", "put your user host")
	var password = flag.String("password", "root", "put your password database")

	// harus menggunakan parse untuk parse semua variable
	// wajib panggil parse
	flag.Parse()

	// * PENGGUNAAN POINTER
	fmt.Println("Host : ", *host)
	fmt.Println("User :", *user)
	fmt.Println("Password : ", *password)
}
