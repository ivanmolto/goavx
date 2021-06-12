package utils

import (
	"bytes"
	"log"
	"net/http"
)

/*
	Helper parameters
*/
type Mintersets struct {
	Minters   []string `json:"minters"`
	Threshold int      `json:"threshold"`
}

type URI struct {
	Address string
	Port    string
}

type Addresses struct {
	Addresses []string `json:"addresses"`
}

type Payload struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  Params `json:"params"`
	ID      int    `json:"id"`
}

type Params struct {
	Assetid    string       `json:"assetID"`
	Payload    string       `json:"payload"`
	To         string       `json:"to"`
	Name       string       `json:"name"`
	Symbol     string       `json:"symbol"`
	Username   string       `json:"username"`
	Password   string       `json:"password"`
	From       []string     `json:"from"`
	Changeaddr string       `json:"changeAddr"`
	Mintersets []Mintersets `json:"minterSets"`
}

/*
  Utility functions
*/
func SendRequest(uri URI, payloadBytes []byte) {
	node := uri.Address + ":" + uri.Port + "/ext/bc/X"
	body := bytes.NewReader(payloadBytes)
	// req, err := http.NewRequest("POST", node, body)
	req, err := http.NewRequest("POST", node, body)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json;")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}

func PayloadValidation(payloadBytes []byte) {

}
