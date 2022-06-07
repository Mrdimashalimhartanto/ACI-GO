package main

import "fmt"

/**
belajar struck function
*/

// type struck untuk customer service
type CustomerService struct {
	Nama, Alamat, JenisKelamin, StatusKaryawan string
	Umur                                       int
}

// type struck untuk detail nasabah
type DetailNasabah struct {
	Name, TempatTinggal, JenisKelamin, StatusNasabah string
	umur                                             int
	jenisPembayaran                                  string
	totalBiaya                                       float32
}

// struck function method
func (customer CustomerService) haiCustomer(namaKaryawan string) {
	fmt.Println("Hallo cuk ", namaKaryawan, "Nama saya ", customer.Nama)
}

// function untuk detail nasabah pemegang polis
func (jenispolis DetailNasabah) haiNasabah(namaNasabah string) {
	fmt.Println("Hallo nasabah ", namaNasabah, "", jenispolis.Name)
}

func main() {

	// make variable to call data yang ada di struck
	var customer CustomerService
	customer.Nama = "Dimas"
	customer.Alamat = "Cipinang"
	customer.JenisKelamin = "laki laki"
	customer.StatusKaryawan = "STAFF"
	customer.Umur = 24
	customer.haiCustomer("Halim")

	var detailPolis DetailNasabah
	detailPolis.Name = "Nani Wijaya Kusuma"
	detailPolis.TempatTinggal = "Kebon Jeruk - Jakarta Barat"
	detailPolis.JenisKelamin = "Perempuan"
	detailPolis.StatusNasabah = "AKTIF"

	detailPolis.umur = 45
	detailPolis.jenisPembayaran = "KARTU CREDIT BRI"
	detailPolis.totalBiaya = 129.000

	// println hasil var detail polis

	fmt.Println(detailPolis.Name)
	fmt.Println(detailPolis.TempatTinggal)
	fmt.Println(detailPolis.JenisKelamin)
	fmt.Println(detailPolis.StatusNasabah)

	fmt.Println(detailPolis.umur)
	fmt.Println(detailPolis.jenisPembayaran)
	fmt.Println(detailPolis.totalBiaya)

	//fmt.Println(customer.Nama,
	//	customer.Alamat,
	//	customer.JenisKelamin,
	//	customer.StatusKaryawan,
	//	customer.Umur)
	//
	//dimas := CustomerService{
	//	Nama:           "Dimas",
	//	Alamat:         "Cipinang",
	//	JenisKelamin:   "Laki Laki",
	//	StatusKaryawan: "STAFF",
	//	Umur:           24,
	//}
	//fmt.Println(dimas)

}
