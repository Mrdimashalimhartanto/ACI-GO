package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	// string interer
	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	//	konversi
	value := strconv.FormatInt(1000000, 10)
	fmt.Println(value)

	// PENGGUNAAN FUNCTION ATIO
	valueInt, _ := strconv.Atoi("14000000")
	fmt.Println(valueInt)
}
