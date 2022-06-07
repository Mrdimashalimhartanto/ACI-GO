package main

import "fmt"

func main() {

	var nilai32 int32 = 128
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Dimas Halim"
	var address = "Jl Cipinang Muara Ilir"
	// BYTE STRING INT8
	// byte di gunakan untuk conversi string
	var e byte = name[0]
	var a byte = address[0]

	var eString string = string(e)
	var aString string = string(a)

	fmt.Println(name)
	fmt.Println(address)
	fmt.Println(eString)
	fmt.Println(aString)

}
