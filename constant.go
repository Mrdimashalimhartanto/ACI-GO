package main

import "fmt"

func main() {

	// MULTIPLE VARIABLE CONSTANTA
	const (
		fullname = "Dimas Halim Hartanto"
		Alamat   = "Cipinang Muara"
		Gawean   = "IT Developer STAFF"
	)

	const (
		JenisKelamin   = "Laki Laki"
		KerjaSampingan = "sembahyang ama silat"
	)

	//const NamaLengkap = "Dimas Halim Hartanto"
	//const Alamat = "Jl Cipinang Muara"

	fmt.Println(fullname)

	fmt.Println(Alamat)
	fmt.Println(Gawean)

	fmt.Println(KerjaSampingan)
	fmt.Println(JenisKelamin)

}
