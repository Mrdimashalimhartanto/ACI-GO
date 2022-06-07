package main

import (
	"fmt"
)

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error pada aplikasi", message)
	}
	fmt.Println("Applikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APPLIKASI ERROR")
	}

	fmt.Println("Applikasi berjalan")
}

func main() {
	runApp(true)
	fmt.Println("Dimas")
}
