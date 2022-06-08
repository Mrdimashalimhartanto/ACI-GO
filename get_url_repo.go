package main

import (
	"belajar-golang-dasar/repository"
	"fmt"
)

func main() {
	//repository.GetURL("Dimas")
	repository.GetURL()
	repository.GetTitle()
	repository.GetHeader()

	fmt.Println(repository.GetURL())
	fmt.Println(repository.GetTitle())
	fmt.Println(repository.GetHeader())
}
