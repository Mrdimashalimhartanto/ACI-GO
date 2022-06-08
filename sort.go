package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type DetailUser struct {
	FullName string
	Umur     int
}

type UserSlice []User
type DetailSlice []DetailUser

func (value DetailSlice) Len() int {
	return len(value)
}

func (value DetailSlice) Less(n, u int) bool {
	return value[u].Umur < value[u].Umur

}

func (value DetailSlice) Swap(n, u int) {
	value[u], value[n] = value[u], value[n]

}

// user slice

func (value UserSlice) Len() int {
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

	details := []DetailUser{
		{"Dimas Halim Hartanto", 24},
		{"Rizky Setiawan", 28},
		{"Ferdi arrahman damin", 29},
	}

	//fmt.Println(users)
	//fmt.Println(details)

	// user slice masuk ke kontrak interface
	sort.Sort(UserSlice(users))
	sort.Sort(DetailSlice(details))
	fmt.Println(users)
	fmt.Println(details)
}
