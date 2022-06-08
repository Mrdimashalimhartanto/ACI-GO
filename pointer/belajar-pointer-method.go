package main

import "fmt"

// buat struct name
type Man struct {
	FullName, Address string
}

type IdentitasNasabah struct {
	NamaLengkap, AlamatTempatTinggal, JenisPolis string
}

type KodeBank struct {
	TextKode                                           string
	KodeBank                                           string
	BCA, MANDIRI, BNI, BRI, BANK, CIMB, NIAGA, SYARIAH string
}

// function pointer method
func (man *Man) Married() {
	man.FullName = "Mr. " + man.FullName + man.Address
}

func (id *IdentitasNasabah) DataNasabah() {
	id.NamaLengkap = "Mr. " + id.NamaLengkap + id.AlamatTempatTinggal + id.JenisPolis
}

func (Kode *KodeBank) DetailKodeBank() {
	Kode.TextKode = "DETAIL KODE BANK BERIKUT :"
	Kode.KodeBank = "KODE : "
}

func main() {
	peserta := Man{"Dimas", " Jl Cipinang"}
	peserta.Married()

	detailinformasi := IdentitasNasabah{"Dimas Halim ", "Jl Cipinang ", "ASURANSI MANDIRI RECEH"}
	detailinformasi.DataNasabah()

	//kodeBank := KodeBank{""}

	fmt.Println(detailinformasi.NamaLengkap)

}
