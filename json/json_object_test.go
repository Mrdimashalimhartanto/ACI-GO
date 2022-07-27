package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Logs struct {
	Title     string
	Action    string
	Type      string
	Request   string
	Response  string
	Message   string
	Order     int
	BrAccount []string
	Addresses []AddressUser
}

type AddressUser struct {
	Street     string
	Country    string
	PostalCode string
}

func TestJSONObject(t *testing.T) {

	datalogs := Logs{
		Title:     "callback in notif ewallet dana",
		Action:    "charge",
		Type:      "In",
		Request:   "ip client 162.158.170.244",
		Response:  "Success",
		Message:   "Success",
		Order:     202200000371,
		BrAccount: []string{"test", "fendy.fendy", "tukang it", "corporate BCA", "tm_ciputralife", "tech team"},

		Addresses: []AddressUser{
			{
				Street:     "Jl Srengseng barat",
				Country:    "Indonesia",
				PostalCode: "13420",
			},
			{
				Street:     "Jl. Jendral Sudirman",
				Country:    "Indonesia",
				PostalCode: "134552",
			},
		},
	}

	bytes, _ := json.Marshal(datalogs)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexCMS(t *testing.T) {
	jsonString := `[{"Street":"Jl Srengseng barat","Country":"Indonesia","PostalCode":"13420"},{"Street":"Jl. Jendral Sudirman","Country":"Indonesia","PostalCode":"134552"}]`
	jsonBytes := []byte(jsonString)

	datalogs := &[]AddressUser{}
	err := json.Unmarshal(jsonBytes, datalogs)
	if err != nil {
		panic(err)
	}
	fmt.Println(datalogs)
}
