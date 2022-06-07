package main

import "fmt"

func main() {

	// TIPE DATA MAP MENGGUNAKAN STRING
	namaNasabah := map[string]string{
		"nama":    "dimas",
		"alamat":  "cipinang",
		"jobdesk": "Programmer",
	}

	MitraBankIndonesia := map[string]string{
		"bank":     "BCA, MANDIRI, BRI, BUKOPIN, BTPN, BANK JAGO, BANK SYARIAH, BANK MANDIRI SYARIAH",
		"kodebank": "340, 123, 224, 212, 33545, 3323,555, 112, 029, 229",
	}

	// bisa di tambahin tanpa di input ke dalam array

	namaNasabah["posisi"] = "staf"

	fmt.Println(namaNasabah)
	fmt.Println(namaNasabah["nama"])
	fmt.Println(namaNasabah["alamat"])
	fmt.Println(namaNasabah["jobdesk"])
	fmt.Println(namaNasabah["posisi"])

	MitraBankIndonesia["bank"] = "BCA, MANDIRI, BRI, BUKOPIN, BTPN, BANK JAGO, BANK SYARIAH, BANK MANDIRI SYARIAH"
	MitraBankIndonesia["kode"] = "340, 123, 224, 212, 33545, 3323,555, 112, 029, 229"

	fmt.Println(MitraBankIndonesia["bank"])
	fmt.Println(MitraBankIndonesia["kode"])

	//	len mengambil jumlah data yang ada di dalam array

	//var NomorPolis map[string]string = make(map[string]string)

	var DetailNasabah map[string]string = make(map[string]string)
	DetailNasabah["nama"] = "Dimas Halim"
	DetailNasabah["jeniskelamin"] = "laki laki"
	DetailNasabah["nopolis"] = "123123123"
	DetailNasabah["alamat"] = "Jakarta Timur"
	fmt.Println(len(DetailNasabah))

	//delete(DetailNasabah, "nopolis")

	//fmt.Println(DetailNasabah)

}
