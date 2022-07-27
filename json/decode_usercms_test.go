package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSONUser(t *testing.T) {
	jsonRequest := `{"FirstName":"Dimas","MiddleName":"Halim","LastName":"Hartanto"}`
	jsonBytes := []byte(jsonRequest)

	users := &Customer{}
	json.Unmarshal(jsonBytes, users)

	fmt.Println(users)
}
