package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONMapRequestDecode(t *testing.T) {
	jsonRequest := `{"order_id":"TE202200000433","action":"charge","type":"Out"}`
	jsonBytes := []byte(jsonRequest)

	var result map[string]interface{}
	_ = json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["order_id"])
	fmt.Println(result["action"])
	fmt.Println(result["type"])
}

func TestJSONMapRequestDataLogsDecode(t *testing.T) {
	jsonRequest := `{"ip_client":"127.0.0.1","headers":{"content-length":["472"],"authorization":["Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJsdW1lbi1qd3QiLCJzdWIiOiJtMnlqbTNlRFVLTkVZWXhNM3Z0WTdtTU9pN3BtV3dtVG5JZmRvNHBZTlMyellqNnkza3NoSllXS2ZtOFd6d21hIiwiaWF0IjoxNjU0OTM5NTYyLCJleHAiOjE2NTQ5NDMxNjIsImJyX2lkIjoxfQ.ZdrAy4mEA3rGemXE1pJ0ndSVdPqDV2wHXpkeG8jRfM8"],"content-type":["application\/x-www-form-urlencoded"],"user-agent":["GuzzleHttp\/7"],"host":["bnepaygcp.ciputralife.com"]},"body":{"paytype":"all","paycode":"all","action":"charge","name":"Ciputralife testing","email":"testing@ciputralife.com","phone":"81319312210","external_id":"ECL2022it433MRs","refno":"","description":{"text_1":"Paket registrasi Ciputra Life-Saver","text_2":"","text_3":""},"amount":"1","exptype":"SD","expval":"2022-07-08","redirect_url":"https:\/\/devebuy.ciputralife.com\/complete\/ECL2022it433MRs","callback_url":"https:\/\/devebuy.ciputralife.com\/api\/v1\/callback","send_email":"0","send_sms":"0"}}`
	jsonBytes := []byte(jsonRequest)

	var result map[string]interface{}
	_ = json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
}

func TestJSONRequestEncode(t *testing.T) {
	detailid := map[string]interface{}{
		"order_id": "TE202200000433",
		"action":   "charge",
		"type":     "Out",
	}
	bytes, _ := json.Marshal(detailid)
	fmt.Println(string(bytes))

}
