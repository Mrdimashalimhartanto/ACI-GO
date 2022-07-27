package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type OrderID struct {
	//TODO: JSON TAG BERFUNGSI UNTUK MERUBAH SEBUAH STRING YANG AWALANNYA PASCAL CASE MENJADI struct case

	OrderId string `json:"order_id"`
	Action  string `json:"action"`
	Type    string `json:"type"`
}

type DataLogs struct {
	Title    string
	OrderId  string
	NoPolis  string
	Action   string
	Type     string
	Request  string
	Response string
}

func TestJSONDataLogs(t *testing.T) {
	hasillogsdata := DataLogs{
		Title:    "request epay success",
		OrderId:  "TE202200000423",
		NoPolis:  "",
		Action:   "charge",
		Type:     "in",
		Request:  "{\"ip_client\":\"127.0.0.1\",\"headers\":{\"content-length\":[\"472\"],\"authorization\":[\"Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJsdW1lbi1qd3QiLCJzdWIiOiJtMnlqbTNlRFVLTkVZWXhNM3Z0WTdtTU9pN3BtV3dtVG5JZmRvNHBZTlMyellqNnkza3NoSllXS2ZtOFd6d21hIiwiaWF0IjoxNjU0OTM5NTYyLCJleHAiOjE2NTQ5NDMxNjIsImJyX2lkIjoxfQ.ZdrAy4mEA3rGemXE1pJ0ndSVdPqDV2wHXpkeG8jRfM8\"],\"content-type\":[\"application\\/x-www-form-urlencoded\"],\"user-agent\":[\"GuzzleHttp\\/7\"],\"host\":[\"bnepaygcp.ciputralife.com\"]},\"body\":{\"paytype\":\"all\",\"paycode\":\"all\",\"action\":\"charge\",\"name\":\"Ciputralife testing\",\"email\":\"testing@ciputralife.com\",\"phone\":\"81319312210\",\"external_id\":\"ECL2022it433MRs\",\"refno\":\"\",\"description\":{\"text_1\":\"Paket registrasi Ciputra Life-Saver\",\"text_2\":\"\",\"text_3\":\"\"},\"amount\":\"1\",\"exptype\":\"SD\",\"expval\":\"2022-07-08\",\"redirect_url\":\"https:\\/\\/devebuy.ciputralife.com\\/complete\\/ECL2022it433MRs\",\"callback_url\":\"https:\\/\\/devebuy.ciputralife.com\\/api\\/v1\\/callback\",\"send_email\":\"0\",\"send_sms\":\"0\"}}",
		Response: "{\"headers\":{},\"original\":{\"code\":\"00\",\"status\":true,\"data\":{\"shortlink\":\"https:\\/\\/clife.id\\/ykDZ6gaE\",\"virtual_account_number\":\"\",\"order_id\":\"TE202200000423\",\"external_id\":\"ECL2022it433MRs\",\"expired_date\":\"2022-07-08\"},\"message\":\"Success\"},\"exception\":null}",
	}

	bytes, _ := json.Marshal(hasillogsdata)
	fmt.Println(string(bytes))
}

func TestJSONTag(t *testing.T) {
	orderid := OrderID{
		OrderId: "TE202200000433",
		Action:  "charge",
		Type:    "Out",
	}

	bytes, _ := json.Marshal(orderid)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"order_id":"TE202200000433","action":"charge","type":"Out"}`
	jsonBytes := []byte(jsonString)

	orderid := &OrderID{}

	json.Unmarshal(jsonBytes, orderid)
	fmt.Println(orderid)
}
