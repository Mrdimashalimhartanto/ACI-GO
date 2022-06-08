package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []User{
		{"Dimas", 40},
		{"Febi", 30},
		{"ferdi", 20},
		{"fajar", 10},
		{"Riski", 60},
	}

	fmt.Println(users)

	// user slice masuk ke kontrak interface
	sort.Sort(UserSlice(users))
	fmt.Println(users)
}
