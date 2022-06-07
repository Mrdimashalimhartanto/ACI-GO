package helper

import "fmt"

var GolangVersion = 112

//var version = 1

func SayHello(name string) {
	fmt.Println("hello", name)
}

func SayGoodBye(name string) {
	fmt.Println("Goodbye", name)
}
