package main

import "fmt"

func main() {

	//TODO: pengambilan string untuk masing masing nilai
	var names [3]string

	names[0] = "Dimas"
	names[1] = "Halim"
	names[2] = "Hartanto"

	//fmt.Println(names[0])
	//fmt.Println(names[1])
	//fmt.Println(names[2])

	var jumlah = [3]int{
		90,
		80,
		60,
	}
	//fmt.Println(jumlah)

	fmt.Println(len(jumlah))
	fmt.Println(len(names))

	var nama = [3]string{
		"Dimas Halim Hartanto",
		"Ferdi damin",
		"Febri sutisna",
	}
	fmt.Println(nama)

	var jeniskelamin = [3]string{
		"Laki Laki",
		"perempuan",
		"laki laki dan perempuan",
	}

	fmt.Println(len(jeniskelamin))

}
