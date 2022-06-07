package main

import "fmt"

func main() {
	//TODO: switch dengan menggunakan variable
	name := "Dimas"
	statuspolis := "sudah bayar"
	detailpolis := "polis sudah aktif"
	statuspembayaran := "melalui BCA"
	jenispembayaran := "CREDIT CARD"

	switch jenispembayaran {
	case "CREDIT CARD":
		fmt.Println("sudah bayar dengan credit card")
	default:
		fmt.Println("asik udah bayar")
	}

	switch statuspembayaran {
	case "melalui BCA":
		fmt.Println("pembayaran mu sudah terkonfirmasi")
	default:
		fmt.Println("terima kasih")
	}

	switch detailpolis {
	case "polis sudah aktif":
		fmt.Println("Hallo Dimas polis mu suda aktif kembali")
		fmt.Println("selamat datang, silahkan login")
	default:
		fmt.Println("silahkan login")
	}

	switch name {
	case "Dimas":
		fmt.Println("Hallo Dimas")
		fmt.Println("pegawai ciputra")
	case "Febi":
		fmt.Println("Hallo febi")
		fmt.Println("pegawai ciputra")
	case "Ferdi":
		fmt.Println("Hallo Ferdi")
		fmt.Println("pegawai ciputra")

	default:
		fmt.Println("Pegawai ciputra not found")
	}

	//TODO: switch dengan perkondisian

	switch statuspolis {
	case "sudah bayar":
		fmt.Println("Hallo Dimas")
		fmt.Println("kamu sudah bayar")
	case "belum bayar":
		fmt.Println("Hallo Dimas")
		fmt.Println("Kamu belum bayar")
	case "Menunggak":
		fmt.Println("Hallo dimas")
		fmt.Println("Kamu masih ada tunggakan nih")

	default:
		fmt.Println("Silahkan cek tagihan mu")
		fmt.Println("silahkan melakukan pembayaran")

	}

	//====================================================================

	// TODO: switch short statement
	switch length := len(name); length > 10 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama Sudah benar")
	}

	//	TODO: switch tanpa kondisi
	length := len(name)
	jumlah := len(statuspolis)

	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("nama sudah benar")
	}

	switch {
	case jumlah > 10:
		fmt.Println("status polis sudah terbayar")
	default:
		fmt.Println("polis sudah terbayar")
	}
}
