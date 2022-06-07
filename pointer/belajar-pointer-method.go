package main

import "fmt"

// buat struct name
type Man struct {
	FullName, Address string
}

type IdentitasNasabah struct {
	NamaLengkap, AlamatTempatTinggal, JenisPolis string
}

// function pointer method
func (man *Man) Married() {
	man.FullName = "Mr. " + man.FullName + man.Address
}

func (id *IdentitasNasabah) DataNasabah() {
	id.NamaLengkap = "Mr. " + id.NamaLengkap + id.AlamatTempatTinggal + id.JenisPolis
}

func main() {
	peserta := Man{"Dimas", " Jl Cipinang"}
	peserta.Married()

	detailinformasi := IdentitasNasabah{"Dimas Halim ", "Jl Cipinang ", "ASURANSI MANDIRI RECEH"}
	detailinformasi.DataNasabah()

	//fmt.Println(peserta.FullName)

	fmt.Println(detailinformasi.NamaLengkap)

}
