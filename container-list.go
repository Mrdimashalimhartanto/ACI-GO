package main

import (
	"container/list"
	"fmt"
)

func main() {

	data := list.New()

	data.PushBack("Dimas")
	data.PushBack("Halim")
	data.PushBack("Hartanto")
	data.PushFront("Adela")

	// pengambilan data depan ke belakang
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	//	pengambilan data belakang ke depan
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}

}
