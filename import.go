package main

import (
	"belajar-golang-dasar/database"
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	helper.SayHello("Dimas")
	helper.SayGoodBye("halim")
	helper.DetailPolis("BCA", "Dimas")
	// todo : import go version
	fmt.Println(helper.GolangVersion)
	fmt.Println(database.StatusConnection())
	fmt.Println(database.IpConnection())

}
