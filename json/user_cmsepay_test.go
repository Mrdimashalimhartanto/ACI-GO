package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Hobbies    []string
	Addresses  []Address
}

func TestUserCmsEPAY(t *testing.T) {
	customer := Customer{
		FirstName:  "Dimas",
		MiddleName: "Halim",
		LastName:   "Hartanto",
		Hobbies:    []string{"Gaming", "Swimming", "Coding"},
		Addresses: []Address{
			{
				Street:     "Jalan Belum jadi",
				Country:    "Indonesia",
				PostalCode: "13420",
			},
			{
				Street:     "Jalan Sudah Jadi",
				Country:    "Indonesia",
				PostalCode: "13522",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecodeUser(t *testing.T) {
	jsonString := `{"FirstName":"Dimas","MiddleName":"Halim","LastName":"Hartanto","Hobbies":["Gaming","Swimming","Coding"],"Addresses":[{"Street":"Jalan Belum jadi","Country":"Indonesia","PostalCode":"13420"},{"Street":"Jalan Sudah Jadi","Country":"Indonesia","PostalCode":"13522"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)

}
