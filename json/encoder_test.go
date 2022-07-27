package main

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("output_data_logs.json")
	encoder := json.NewEncoder(writer)

	datalogs := Logs{
		Title:     "callback in notif ewallet dana",
		Action:    "charge",
		Type:      "In",
		Request:   "ip client 162.158.170.244",
		Response:  "Success",
		Message:   " Success",
		Order:     202200000371,
		BrAccount: []string{"test", "fendy.fendy", "tukang IT", "Corporate BCA", "tm_ciputralife", "tech team"},
		Addresses: []AddressUser{
			{
				Street:     "Jl Srengseng barat",
				Country:    "Indonesia",
				PostalCode: "13420",
			},
			{
				Street:     "Jl Jendral Sudirman",
				Country:    "Indonesia",
				PostalCode: "1345335",
			},
		},
	}

	_ = encoder.Encode(datalogs)

}
