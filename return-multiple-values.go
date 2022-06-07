package main

import "fmt"

// TODO: belajar function return multiple value

// TODO: INFONASABAH
func InfoNasabah() (string, string, string, string) {
	return "Dimas", "Halim", "IT Staff", "Ciputra Life"
}

// todo : alamat tinggal nasabah
func AlamatTinggalNasabah() (string, string, string) {
	return "Jl.Cipinang Muara", "RT 14 / 08", "Kelurahan Cipinang Muara"
}

func Bank() (string, string, string) {
	return "BCA", "MANDIRI", "BRI"
}

func main() {
	namaDepan, _, jabatanKerja, tempatKerja := InfoNasabah()
	fmt.Println(namaDepan, jabatanKerja, tempatKerja)

	//	todo : alamat tempat tinggal
	AlamatNasabah, RTNasabah, KelurahanNasabah := AlamatTinggalNasabah()
	fmt.Println(AlamatNasabah, RTNasabah, KelurahanNasabah)

	kartuBCA, kartuMandiri, kartuBRI := Bank()
	fmt.Println(kartuBCA, kartuMandiri, kartuBRI)

}
