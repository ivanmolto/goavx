package assets

import (
	"encoding/json"
	"log"

	. "github.com/smartpassnft/goavx/avm/utils"
)

type NFTPayload struct {
	Jsonrpc string    `json:"jsonrpc"`
	Method  string    `json:"method"`
	Params  nftParams `json:"params"`
	ID      int       `json:"id"`
}

type NFTUTXOPayload struct {
	Jsonrpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
	ID      int      `json:"id"`
}

type nftParams struct {
	Name       string       `json:"name"`
	Symbol     string       `json:"symbol"`
	Mintersets []Mintersets `json:"minterSet"`
	Username   string       `json:"username"`
	Password   string       `json:"password"`
}

// TODO: Create an account with balance to allow users to mint NFT's
func CreateNFTAsset(data Payload, uri URI) {
	nftAsset := NFTPayload{
		Jsonrpc: data.Jsonrpc,
		ID:      data.ID,
		Method:  data.Method,
		Params: nftParams{
			Name:       data.Params.Name,
			Symbol:     data.Params.Symbol,
			Mintersets: data.Params.Mintersets,
			Username:   data.Params.Username,
			Password:   data.Params.Password,
		},
	}

	payloadBytes, err := json.Marshal(nftAsset)
	if err != nil {
		log.Fatal(err)
	}

	SendRequest(uri, payloadBytes)
}

//Mint NFT
func MintNFTAsset(data Payload, uri URI) {
	// TODO: Check payload params for required params
	// TODO: Form execution and call from createNFTAsset

}

// Send NFT
func SendNFT(data Payload, uri URI) {
	// TODO: Check payload params for required params
	// TODO: Check if users owns NFT

}

func GetUTXOS(data Payload, address []string, uri URI) {
	utxo := NFTUTXOPayload{
		Jsonrpc: data.Jsonrpc,
		ID:      data.ID,
		Method:  data.Method,
		Params:  address,
	}

	payloadBytes, err := json.Marshal(utxo)
	if err != nil {
		log.Fatal(err)
	}
	SendRequest(uri, payloadBytes)
}
