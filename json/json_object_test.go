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
	Addresses []AccountBroker
}

type AccountBroker struct {
	Email   string
	Name    string
	Address string
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

		Addresses: []AccountBroker{
			{
				Email:   "test@ciputralife.com",
				Name:    "test",
				Address: "Ciputra Life",
			},
			{
				Email:   "fendy.fendy@ciputralife.com",
				Name:    "Fendy",
				Address: "Jakarta",
			},
		},
	}

	bytes, _ := json.Marshal(datalogs)
	fmt.Println(string(bytes))
}
