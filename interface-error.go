package main

import (
	"errors"
	"fmt"
)

// error interface sudah ada didalam golang.cukup panggil kontrak default dari golang
// function error interface

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

// todo: function untuk tambah
func Tambah(total int, tambah int) (int, error) {
	if tambah == 0 {
		return 0, errors.New("jangan banyak banyak nambahanya")
	} else {
		result := total + tambah
		return result, nil
	}
}

// todo: function untuk perkalian
func Perkalian(total int, kali int) (int, error) {
	if kali == 0 {
		return 0, errors.New("Jangan kebanyakan kali")
		//result := total * kali
	} else {
		result := total * kali
		return result, nil
	}
}

func main() {
	// hasil pembagi
	hasil, err := Pembagi(100, 10)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}

	// hasil pertambahan
	hasiltambah, err := Tambah(200, 200)
	if err == nil {
		fmt.Println("Hasil", hasiltambah)
	} else {
		fmt.Println("Error", err.Error())
	}

	//	hasil perkalian
	hasilPerkalian, err := Perkalian(400, 500)
	if err == nil {
		fmt.Println("Total", hasilPerkalian)
	} else {
		fmt.Println("Error", err.Error())
	}
}
