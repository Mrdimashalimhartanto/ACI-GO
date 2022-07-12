package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("d([a-z])s")

	fmt.Println(regex.MatchString("dis"))
	fmt.Println(regex.MatchString("dts"))
	fmt.Println(regex.MatchString("dDs"))
	fmt.Println(regex.MatchString("dbs"))

	//regex.FindAllString("dimas ")
	cari := regex.FindAllString("dis das dus dos des dds dks dbs", 6)
	fmt.Println(cari)
}
