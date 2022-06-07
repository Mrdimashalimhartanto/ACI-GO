package main

import "fmt"

func main() {

	//TODO: perulangan dengan for
	///*counter := 1
	//
	//for counter <= 10 {
	//	fmt.Println("Perulangan", counter)
	//	counter++
	//}*/

	//TODO: kode program for dengan statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("perulangan ke", counter)
	}

	//	TODO: FOR RANGE
	slice := []string{"Dimas", "Halim", "Hartanto", "Adela", "Fauziyah"}

	// TODO: cara manual mengambil array
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//	TODO : cara for range array slice berupa ( index )
	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	// TODO: for range dengan map ( menggunakan key dan value )
	person := make(map[string]string)
	person["name"] = "Dimas"
	person["title"] = "IT Staff"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}

	//	TODO: PAKAI UNDERSCORE _ UNTUK MENANDAKAN BAHWA VARIABLE TIDAK DI PAKAI

}
