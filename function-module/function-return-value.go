package main

import "fmt"

func getNamaNasabah(nama string) string {
	
	//TODO: NILAI FUNCTION MENGGUNAKAN PERKONDISIAN

	if nama == "" {
		return "Hallo nasabah"
	} else {
		return "Hello " + nama
	}

}

func main() {
	result := getNamaNasabah("Dimas")
	fmt.Println(result)

	fmt.Println(getNamaNasabah(""))
}
