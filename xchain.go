package govalanche

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

/*
	Function Specific Payloads
*/
type Payload struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  Params `json:"params"`
	ID      int    `json:"id"`
}

type NFTPayload struct {
	Jsonrpc string    `json:"jsonrpc"`
	Method  string    `json:"method"`
	Params  nftParams `json:"params"`
	ID      int       `json:"id"`
}

type UserPayload struct {
	Jsonrpc string     `json:"jsonrpc"`
	Method  string     `json:"method"`
	Params  UserParams `json:"params"`
	ID      int        `json:"id"`
}

type NFTUTXOPayload struct {
	Jsonrpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
	ID      int      `json:"id"`
}

/*
	Function Specific Parameters for JSON RPC
*/
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

type UserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type nftParams struct {
	Name       string       `json:"name"`
	Symbol     string       `json:"symbol"`
	Mintersets []Mintersets `json:"minterSet"`
	Username   string       `json:"username"`
	Password   string       `json:"password"`
}

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

// helper functions to parse through multiple UTXO's on account
func AssetUTXO() {

}

/*
  X-Chain NFT Functions

  Current implemented functions
  - CreateNFTAsset
  - MintNFTAsset
  - SendNFT
*/

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

/*
  User related functions

  Current implemented functions
  - getAddress
  - CreateUser
  - CreateAddress
*/

// Get X-Chain Address
func GetAddress(data Payload, uri URI, seed []byte) *bip32.Key {
	// TODO: Pass string as string and convert to []byte
	if seed != nil {
		if bip39.IsMnemonicValid(string(seed[:])) {
			masterKey, _ := bip32.NewMasterKey(seed)
			publicKey := masterKey.PublicKey()
			return publicKey
		}
	}
	return nil
	// TODO: Get address from user password combination
}

// CreateUser function
func CreateUser(data Payload, uri URI) {
	// Create address using json request
	user := UserPayload{
		Jsonrpc: data.Jsonrpc,
		Method:  data.Method,
		ID:      data.ID,
		Params: UserParams{
			Username: data.Params.Username,
			Password: data.Params.Password,
		},
	}

	payloadBytes, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	SendRequest(uri, payloadBytes)
}

// Generate Address from mmemonic
func GenerateAddress(pass string, num int) (string, *bip32.Key) {
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	seed := bip39.NewSeed(mnemonic, pass)
	masterKey, _ := bip32.NewMasterKey(seed)
	publicKey := masterKey.PublicKey()

	return string(seed[:]), publicKey
}

// Create user Address
func CreateAddressAccount(data Payload, uri URI) {

	// Create address using json request
	user := UserPayload{
		Jsonrpc: data.Jsonrpc,
		Method:  data.Method,
		ID:      data.ID,
		Params: UserParams{
			Username: data.Params.Username,
			Password: data.Params.Password,
		},
	}

	payloadBytes, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	SendRequest(uri, payloadBytes)
}

func CreateAddressSeed(seed []byte) *bip32.Key {
	// TODO: Validate seed mnemonic
	if bip39.IsMnemonicValid(string(seed[:])) {
		masterKey, _ := bip32.NewMasterKey(seed)
		publicKey := masterKey.PublicKey()
		// Return key generated from address
		return publicKey
	} else {
		log.Fatal("Invalid Mnemonic")
	}
	return nil
}
