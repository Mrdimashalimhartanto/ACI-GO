package main

import "fmt"

/**
pass by value. secara default di golang semua variable itu di passing by value
bukan by refrence

KETERANGAN
1 .pass by value dalam pointer golang bekerja untuk melakukan copy data pada value
2. tanda & adalah sebuah tanda pointer dari sebuah value. selain itu,
tanda bahwa value ada sebuah pointer bisa di tandai dengan BINTANG ( * )
3. tanda bintang * bisa juga di gunakan sebagai pointer bisa juga di gunakan sebagai operator
4. jika kita mau membuat sebuah pointer baru cukup di tambahkan *pointer kita + & untuk membuat data baru
5. Jika kita mau membuat pointer baru cukup dengan menggunakan new
6. tanda * bisa digunakan sebagai penanda pointer pada sebuah nilai NILAI
7. pointer function digunakan untuk membuat sebuah pointer didalam function
*/

// pembuatan struck untuk alamat
type Address struct {
	City, Province, Country string
}

type PaymentGateway struct {
	NamePG, DetailPG, MitraPG string
	NoPG                      int
	DetailPaymentPG           string
}

// buat function untuk payment gateway
func PaymentGatewayMethod(pg_detail *PaymentGateway) {
	pg_detail.NamePG = "BCA"
}

// pointer function
func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

// struck untuk nama lengkap
type InformasiNasabah struct {
	NamaLengkap, Alamat, TempatLahir, Kelurahan, kota string
	umur                                              int
	penghasilan                                       int
}

// struck untuk produck asuransi
type ProductAsuransi struct {
	NamaProduk, manfaat, caraMembayar string
	harga                             int
}

// pointer cms-epay
type CmsEpay struct {
	JenisPembayaran, JenisKartu string
	Kadarluwarsa                int
}

func main() {

	CmsCIA1 := CmsEpay{"Tunai", "Credit Card", 2022}
	CmsCIA2 := &CmsCIA1
	CmsCIA3 := &CmsCIA1

	fmt.Println(CmsCIA2)
	fmt.Println(CmsCIA3)

	CmsCIA2.JenisKartu = "BCA Debit Card"
	// POINTER
	*CmsCIA2 = CmsEpay{"Non Tunai", "BCA Syariah", 2026}
	fmt.Println(CmsCIA2)

	// new digunakan untuk membuat sebuah pointer baru
	var CmsCIA4 *CmsEpay = new(CmsEpay)
	CmsCIA4.JenisKartu = "BRI PRATAMA PLUS PLUS"
	fmt.Println(CmsCIA4)

	// pointer address ===========================================================
	address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	address2 := &address1
	address2.City = "Lampung"

	fmt.Println(address1)
	fmt.Println(address2)

	// ========== STRUCK UNTUK INFORMASI NASABAH
	br_nasabah := InformasiNasabah{"Dimas Halim Hartanto", "Jl.Cipinang Muara", "Jakarta", "Cipinang Muara",
		"Jakarta Timur", 24, 5800000}
	br_nasabah2 := &br_nasabah

	fmt.Println(br_nasabah)
	fmt.Println(br_nasabah2)

	// ========== STRUCK UNTUK INFORMASI PAYMENT GATEWWAY

	// ========== STRUCK product asuransi jiwa ciputra
	prd_asuransi := ProductAsuransi{"Asuransi Jiwa Ciputra Mandiri", "test 100 %", "dengan kartu kredit", 140000}
	prd_asuransi2 := &prd_asuransi

	fmt.Println(prd_asuransi)
	fmt.Println(prd_asuransi2)

	// pointer new digunakan untuk membuat sebuah pointer baru
	var Alamat1 *Address = new(Address)
	fmt.Println(Alamat1)

	// pointer  function
	var alamat = Address{
		City:     "Jakarta",
		Province: "Sumatra",
		Country:  "",
	}
	var alamatPointer *Address = &alamat
	ChangeCountryToIndonesia(alamatPointer)
	fmt.Println(alamat)

	var detail_pg = PaymentGateway{
		NamePG:          "BCA",
		DetailPG:        "pembayaran melalui midrtans",
		MitraPG:         "Midtrans",
		NoPG:            2740603354,
		DetailPaymentPG: "Pembayaran Polis Asuransi",
	}

	var detailPG_pointer *PaymentGateway = &detail_pg
	PaymentGatewayMethod(detailPG_pointer)
	fmt.Println(detail_pg)

}
