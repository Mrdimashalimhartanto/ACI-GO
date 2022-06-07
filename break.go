package main

import "fmt"

func main() {

	// TODO: pengulangan menggunakan break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("perulangan", i)
	}
}
