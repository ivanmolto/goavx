package goavx

import (
	"encoding/json"
	"log"

	_ "github.com/smartpassnft/goavx/utils"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

/*
	Function Specific Payloads
*/
type UserPayload struct {
	Jsonrpc string     `json:"jsonrpc"`
	Method  string     `json:"method"`
	Params  UserParams `json:"params"`
	ID      int        `json:"id"`
}

/*
	Function Specific Parameters for JSON RPC
*/
type UserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
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
