package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2020, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	//TODO format time rdfe3339

	// FORMAT BAWAAN GOLANG
	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "1990-02-03")
	fmt.Println(parse)

}
