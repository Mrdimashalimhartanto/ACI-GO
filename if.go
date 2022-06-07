package main

import "fmt"

func main() {

	name := "Dimas"

	if name == "Dimas" {
		fmt.Println("Dimas employee")
	} else if name == "Lisa" {
		fmt.Println("Hello Lisa")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("employee not found")
	}

	//	if short statement
	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}

}
