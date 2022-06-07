package main

import "fmt"

func notifikasiNasabah(detailNotif string) string {
	return "Hallo terima kasih sudah melakukan pembayaran" + detailNotif
}

func websiteCallCenter(halamanCS string, kontakcs string) string {
	return "silahkan buka www.ciputralife.com" + halamanCS + kontakcs
}

func ConnectionDatabase(ConnectionName string, DatabaseName string) string {
	return "Connection : " + ConnectionName + DatabaseName
}

func main() {
	// TODO : VARIABLE BISA DI MASUKAN KE DALAM FUNCTION

	infonotif := notifikasiNasabah
	result := infonotif(" terima kasih")
	fmt.Println(result)

	websitecs := websiteCallCenter
	hasil := websitecs(" silahkan klik link ", " / hub kontak ini 123123")
	fmt.Println(hasil)

	infoConnection := ConnectionDatabase
	hasilConnection := infoConnection("localhost ", " db_eservicing ")
	fmt.Println(hasilConnection)

}
