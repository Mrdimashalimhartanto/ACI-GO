package main

import (
	"fmt"
)

func main() {

	//HARUS MENGGUNAKAN VAR UNTUK MEMBUAT VARIABLE MEMANGGILA VARIABLE YANG SUDAH KITA BUAT
	var name string
	var address string
	var sex string
	var job string
	var umur int

	name = "Dimas Halim Hartanto"
	fmt.Println(name)

	address = "Jalan Cipinang Muara"
	fmt.Println(address)

	sex = "Laki - Laki"
	fmt.Println(sex)

	job = "IT Staff"
	fmt.Println(job)

	umur = 40
	fmt.Println(umur)

	//var age = 30
	//fmt.Println(age)

	//	tanpa menyebut var
	var pengalamakerja = "BRI Danareksa Sekuritas"
	fmt.Println(pengalamakerja)

	// PENGGANTI KATA STRING BISA PAKAI :=
	negara := "Indonesia"
	fmt.Println(negara)

	//	membuat multiple variable
	var (
		Namalengkap = "Dimas Halim Hartanto"
		Alamat      = "Jalan Cipinang Muara"
	)

	fmt.Println(Namalengkap)
	fmt.Println(Alamat)
}
