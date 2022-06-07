package main

import "fmt"

func main() {

	var name1 = "Dimas"
	var name2 = "Alim"

	var result bool = name1 > name2
	fmt.Println(result)

	var value1 = 200
	var value2 = 400

	//operation modular
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
