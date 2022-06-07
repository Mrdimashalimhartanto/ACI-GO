package main

import "fmt"

type HasName interface {
	GetName() string
}
type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func SayHello(hashName HasName) {
	fmt.Println("Hello", hashName.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var dimas Person
	dimas.Name = "Dimas"

	SayHello(dimas)

	kuda := Animal{
		Name: "Kuda lumping",
	}
	SayHello(kuda)
}
