package main

import "fmt"

func namaNasabah(namaDepan string, namaBelakang string, jenisKelamin string) {
	fmt.Println("Hallo", namaDepan, namaBelakang, jenisKelamin)
}

func NamaBank(namaDepanBank string, kodeBank int) {
	fmt.Println("Detail Bank ", namaDepanBank, kodeBank)
}

func KodePos(TempatPost string, kode int) {
	fmt.Println("POS INDONESIA", 043)
}

func main() {
	namaDepan := "Dimas"
	namaDepanBank := "Detail Bank"
	kodepost := "Kode pos"

	namaNasabah(namaDepan, "Halim", "laki laki")
	NamaBank(namaDepanBank, 034)
	KodePos(kodepost, 034)
	//namaNasabah("Ferdi", "Damin", "laki laki")

}
