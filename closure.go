package main

import "fmt"

func main() {

	name := "Dimas"
	address := "Jakarta"
	counter := 0

	increment := func() {
		name := "Halim"
		address := "Beda alamat"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
		fmt.Println(address)
	}

	increment()
	increment()

	fmt.Println(name)
	fmt.Println(address)
}
