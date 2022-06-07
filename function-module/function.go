package main

import "fmt"

// TODO : FUNCTION bisa membuat function sendiri contoh nya pada flutter bisa membuat widget sendiri

func detailNasabah() {
	//fmt.Println("Hallo nasabah")

	const (
		fullname = "Dimas"
		address  = "cipinang"
		status   = "Active Polis"
	)

	fmt.Println(fullname)
	fmt.Println(address)
	fmt.Println(status)
}

func statusNasabah() {
	const (
		title           = "IT Developer"
		degree          = "S1"
		gradeofyear int = 2021
	)

	fmt.Println(title)
	fmt.Println(degree)
	fmt.Println(gradeofyear)
}

func nomorPolis() {
	const (
		idPolis      int = 3175030212970003
		tanggalLahir int = 19971202
		idNasabah    int = 10
	)

	fmt.Println(idPolis)
	fmt.Println(tanggalLahir)
	fmt.Println(idNasabah)
}

func main() {
	detailNasabah()
	statusNasabah()
	nomorPolis()
}
