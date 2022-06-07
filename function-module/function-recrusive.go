package main

import "fmt"

// TODO: RECRUSIVE ADALAH SEBUAH METODE YANG DIGUNAKAN UNTUK MELAKUKAN PERKALIAN TERHADAP SEBUAH BILANGAN FUNCTION

func factorialRecrusive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecrusive(value-1)
	}
}

func main() {
	recrusive := factorialRecrusive(10)
	fmt.Println(recrusive)
}
