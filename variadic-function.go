package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {

	total := sumAll(10, 90, 20, 30, 40, 50)
	fmt.Println(total)

	// TODO: SLICE variadic function
	// TODO: JIKA ADA DATA SLICE WAJIB MENGGUNAKAN TANDA [] untuk menentukan berapa jumlah data slice
	slice := []int{10, 20, 30, 40, 50, 60}

	total = sumAll(slice...)

	fmt.Println(total)
}
