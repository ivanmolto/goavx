package avm

import (
	"encoding/json"
	"log"

	"github.com/smartpassnft/goavx/avm/utils"
	bip32 "github.com/tyler-smith/go-bip32"
	bip39 "github.com/tyler-smith/go-bip39"
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

/*
  User related functions

  Current implemented functions
  - getAddress
  - CreateUser
  - CreateAddress
*/

// Get X-Chain Address
func GetAddress(data utils.Payload, uri utils.URI, seed []byte) *bip32.Key {
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
func CreateUser(data utils.Payload, uri utils.URI) {
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
	utils.SendRequest(uri, payloadBytes)
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
func CreateAddressAccount(data utils.Payload, uri utils.URI) {

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
	utils.SendRequest(uri, payloadBytes)
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
