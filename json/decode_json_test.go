package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonString := `{"Title":"callback in notif ewallet dana","Action":"charge","Type":"In","Request":"ip client 162.158.170.244","Response":"Success","Message":"Success","Order":202200000371}`
	jsonBytes := []byte(jsonString)

	datalogs := &Logs{}

	err := json.Unmarshal(jsonBytes, datalogs)
	if err != nil {
		panic(err)
	}

	fmt.Println(datalogs)
	fmt.Println(datalogs.Title)
	fmt.Println(datalogs.Action)
	fmt.Println(datalogs.Request)
	fmt.Println(datalogs.Message)
}

func TestDatalogsDecode(t *testing.T) {
	jsonString := `{"Title":"callback in notif ewallet dana","Action":"charge","Type":"In",
"Request":"ip client 162.158.170.244","Response":"Success","Message":"Success","Order":202200000371,"BrAccount":["test","fendy.fendy","tukang it","corporate BCA","tm_ciputralife","tech team"]}`
	jsonBytes := []byte(jsonString)

	dataaccountcl := &Logs{}
	//TODO: LEBIH BAIK MENGGUNAKAN POINTER UNTUK SETIAP MENGAKSES DATA YANG ADA DI DALAM SLICE ARRAY

	err := json.Unmarshal(jsonBytes, dataaccountcl)
	if err != nil {
		panic(err)
	}

	fmt.Println(dataaccountcl)
	fmt.Println(dataaccountcl.BrAccount)
	fmt.Println(dataaccountcl.Response)
	fmt.Println(dataaccountcl.Addresses)
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"Title":"callback in notif ewallet dana","Action":"charge","Type":"In","Request":"ip client 162.158.170.244","Response":"Success","Message":"Success","Order":202200000371,"BrAccount":["test","fendy.fendy","tukang it","corporate BCA","tm_ciputralife","tech team"],"Addresses":[{"Email":"test@ciputralife.com","Name":"test","Address":"Ciputra Life"},{"Email":"fendy.fendy@ciputralife.com","Name":"Fendy","Address":"Jakarta"}]}`
	jsonBytes := []byte(jsonString)

	customer := &AccountBroker{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Name)
	fmt.Println(customer.Address)
}
