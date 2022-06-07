package main

/**
blank identifier di gunakan untuk tidak menggunakan function.cukup pakai tanda underscore _
*/

// jika tidak mau import func GetDatabase lebih baik menggunakan tanda _ / identifier
import (
	"belajar-golang-dasar/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
