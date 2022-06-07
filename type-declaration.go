package main

import "fmt"

func main() {
	type NoKTP string
	type StatusNikah bool
	type umur int
	type AlamatTinggal string
	type KerjaSampingan string

	//	variable tipe data alias
	var noKTPDimas NoKTP = "3175030212970003"
	var statusnikah StatusNikah = false
	var umurPasangan umur = 24
	var alamatTinggal AlamatTinggal = "Jl. Cipinang Muara"
	var kerjaansampingan KerjaSampingan = "Silat"

	fmt.Println(noKTPDimas)
	fmt.Println(statusnikah)
	fmt.Println(umurPasangan)
	fmt.Println(alamatTinggal)
	fmt.Println(kerjaansampingan)
}
