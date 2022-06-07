package main

import (
	"fmt"
)

type BlackList func(string) bool
type RegisterList func(string) bool
type MasaBerlakuKartu func(string) bool
type KartuIdentitas func(string) bool

func tipekartu(name string, detailKartuNasabah KartuIdentitas) {
	if detailKartuNasabah(name) {
		fmt.Println("kartu kamu adalah : ", name)
	} else {
		fmt.Println("hello", name)
	}
}

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("kamu di block", name)
	} else {
		fmt.Println("hallo dimas", name)
	}
}

func halamanRegis(name string, detailRegister RegisterList) {
	if detailRegister(name) {
		fmt.Println(name)
	} else {
		fmt.Println("Kartu mu di blokir", name)
	}
}

func masaberlakukartu(name string, masaberlaku MasaBerlakuKartu) {
	if masaberlaku(name) {
		fmt.Println("Upps kartu mu masa berlaku nya mau habis ni ", name)
	} else {
		fmt.Println("Kartu mu sudah berlaku lagi ", name)
	}
}

func main() {
	detailRegister := func(name string) bool {
		return name == "Nasabah"
	}

	halamanRegis("Hallo", detailRegister)

	//==================================================================\\
	blacklist := func(name string) bool {
		return name == "dimas"
	}

	detailkartuNasabah := func(name string) bool {
		return name == "Kartu nasabah"
	}

	masaberlaku := func(name string) bool {
		return name == "Dimas"
	}

	masaberlakukartu("yeay kartu mu sudah bisa", masaberlaku)
	registerUser("admin", blacklist)
	registerUser("dimas", blacklist)
	tipekartu("kartu sudah ketemu", detailkartuNasabah)
}
