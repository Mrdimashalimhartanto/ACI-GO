package main

import "fmt"

func main() {

	var months = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var DataBank = [...]string{
		"BCA",
		"MANDIRI",
		"BUKOPIN",
		"BANK MUAMALAT",
		"BCA SYARIAH",
		"BNI",
		"BTPN",
		"BANK DKI",
		"BANK JAGO",
	}

	var nopolis = [...]string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
	}

	var polisslice = nopolis[4:5]
	fmt.Println(polisslice)

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// index slice 2
	var slice2 = months[11:]
	fmt.Println(slice2)

	//function data append
	var slice3 = append(slice2, "Test")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	// function slice bank
	var bankslice = DataBank[1:9]
	fmt.Println(bankslice)
	fmt.Println(len(DataBank))

	//	function make tipe slice
	newSlice := make([]string, 2, 5)

	newSlice[0] = "Dimas"
	newSlice[1] = "Halim"

	fmt.Println(newSlice)
	fmt.Println(cap(newSlice))
	fmt.Println(len(newSlice))

	newDataPolis := make([]string, 3, 5)

	newDataPolis[0] = "11122534995"
	newDataPolis[1] = "23400495954"
	newDataPolis[2] = "32352354334"

	//function len dan cap
	fmt.Println(newDataPolis)
	fmt.Println(len(newDataPolis))
	fmt.Println(cap(newDataPolis))

	copyData := make([]string, len(newDataPolis), cap(newDataPolis))
	copy(copyData, newDataPolis)
	fmt.Println(copyData)

	//	mak function detail info nasabah
	newDataNasabah := make([]string, 4, 5)

	newDataNasabah[0] = "Dimas Halim Hartanto"
	newDataNasabah[1] = "Jl Cipinang Muara"
	newDataNasabah[2] = "IT Developer"

	fmt.Println(newDataNasabah)
	fmt.Println(len(newDataNasabah))
	fmt.Println(cap(newDataNasabah))

	// function copy
	copyDataNasabah := make([]string, len(newDataNasabah), cap(newDataNasabah))
	copy(copyDataNasabah, newDataNasabah)
	fmt.Println(copyDataNasabah)

}
