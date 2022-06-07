package main

import "fmt"

// PENGGUNAAN INTERFACE KOSONG

func UpsSalah(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups Salah"
	}
}

func main() {
	var data interface{} = UpsSalah(2)
	fmt.Println(data)
}
