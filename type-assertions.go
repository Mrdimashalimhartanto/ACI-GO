package main

import "fmt"

/**
type assertion merupakan kemampuan merubah tipe data menjadi
tipe data yang di inginkan
*/

func random() interface{} {
	// kalo pake boolean akan terjadi PANIC
	return "ups ada dimas"
}

func main() {
	//var result interface{} = random()
	//var resultString string = result.(string)
	//fmt.Println(resultString)

	//	Kode Program Type Assertions Switch
	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println("Waduh", value, "is string")
	case int:
		fmt.Println("Int", value, "is integer")
	default:
		fmt.Println("Unknown type")
	}
}
