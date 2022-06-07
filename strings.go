package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Contains("Dimas Halim", "Dimas"))

	fmt.Println(strings.Contains("Dimas Halim", "Kusuma"))
	fmt.Println(strings.Split("Dimas Halim Hartanto", " "))

	fmt.Println(strings.ToLower("DIMAS HALIM HARTANTO"))
	fmt.Println(strings.ToUpper("DIMAS HALIM HARTANTO"))

	fmt.Println(strings.Trim("   Dimas Halim    ", "    "))
	fmt.Println(strings.ReplaceAll("Dimas Dimas Dimas", "Dimas", "Adel"))
}
