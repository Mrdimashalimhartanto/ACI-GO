package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

}

func TestEncode(t *testing.T) {
	logJson("Title")
	logJson(1)
	logJson(true)
	logJson([]string{"callback in notif ewallet dana", "TE202200000371", "charge"})
	logJson("Title")
	logJson(2)
	logJson(true)
	logJson([]string{"update job expired all", "TE202200000415", "charge"})
}
