package main

import "fmt"

func logging() {
	fmt.Println("selesai panggil function")
}

func DetailPolis() {
	fmt.Println("Hai nasabah ini detail polis kamu")
}

func runApplication(value int) {
	// todo : deffer harus di atas berfungsi untuk memanggil function
	defer logging()
	defer DetailPolis()
	fmt.Println("Run App")
	result := 10 / value
	fmt.Println("Result", result)
}

func main() {
	runApplication(10)
	DetailPolis()
}
