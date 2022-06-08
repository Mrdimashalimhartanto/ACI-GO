package main

import "fmt"

func infoNasabah() (namaDepan string, namaTengah string, namaBelakang string) {
	namaDepan = "Dimas"
	namaTengah = "Halim"
	namaBelakang = "Hartanto"
	return
}

func detailPolis() (jenisPolis string, kartuPolis string, alamatPolis string) {
	jenisPolis = "POLIS ASURANSI 1"
	kartuPolis = "BCA"
	alamatPolis = "Jakarta"

	return
}

func Broker() (NamaBroker string, alamatBroker string, statusBroker string) {
	NamaBroker = "PT. ASURANSI CIPUTRA LIFE"
	alamatBroker = "Jl.Iskandar muda selatan, Banten, Jawa Barat"
	statusBroker = "ACTIVE MEMBER"

	return

}

func InformasiVendor() (NamaVendor string, AlamatVendor string) {
	NamaVendor = "PT XL Asuransi Life"
	AlamatVendor = "jl. patiunus 3"

	return

}

func main() {
	namaDepan, namaTengah, namaBelakang := infoNasabah()
	fmt.Println(namaDepan, namaTengah, namaBelakang)

	jenisPolis, kartuPolis, alamatPolis := detailPolis()
	fmt.Println(jenisPolis, kartuPolis, alamatPolis)

	NamaBroker, namaTengah, namaBelakang := Broker()
	fmt.Println(NamaBroker, namaTengah, namaBelakang)

	NamaVendor, AlamatVendor := InformasiVendor()
	fmt.Println(NamaVendor, AlamatVendor)

	//	konstanta untuk masing masing variable
	const (
		fullName = "Dimas"
	)
	fmt.Println(fullName)
}
