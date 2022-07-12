package main

import (
	"context"
	"fmt"
	"testing"
)

//TODO: membuat go context

/**
TODO: go context di gunakan untuk mengirim data request atau sinyal ke proses lain context adalah sebuah interface
*/

//TODO: func context bisa menggunakan background / todo

func TestContext(t *testing.T) {
	backgorund := context.Background()
	fmt.Println(backgorund)

	todo := context.TODO()
	fmt.Println(todo)
}

// TODO: parent & child context & Test Context With Value

//TODO: context with value
func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	//TODO: DATA CONTEXT ITU IMUTABLE = DATA TIDAK BISA DI RUBAH

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")
	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	//fmt.Println(contextF.Value("b"))

}
