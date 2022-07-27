package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestDecoder(t *testing.T) {
	reader, _ := os.Open("Datalogs.json")
	decoder := json.NewDecoder(reader)

	logsdata := &Logs{}
	decoder.Decode(logsdata)
	//_ = decoder.Decode(&logsdata)

	fmt.Println(logsdata)

}
