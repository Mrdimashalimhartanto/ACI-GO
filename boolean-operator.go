package main

import "fmt"

func main() {

	var ujian = 80
	var absensi = 75

	var lulusUjian = ujian >= 80
	var lulusAbsensi = absensi >= 80
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = (lulusUjian && lulusAbsensi)
	fmt.Println(lulus)

	var nilaisks = 90
	var absensimasuk = 80

	//var nilaimahasiswa = nilaisks >= 90
	//var nilailulus = absensimasuk >= 70
	//
	//var kelulusan = (nilaimahasiswa && nilailulus)
	//
	//fmt.Println(kelulusan)

	fmt.Println(nilaisks >= 90 && absensimasuk >= 80)

	//fmt.Println(ujian >= 80 && absensi >= 70)

}
