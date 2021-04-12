package govalanche

import (
  // "log"
  "net/http"
  "bytes"
  "encoding/json"
  "github.com/tyler-smith/go-bip39"
  "github.com/tyler-smith/go-bip32"
)

type Payload struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  Params `json:"params"`
	ID      int    `json:"id"`
}

type Mintersets struct {
  Minters   []string  `json:"minters"`
  Threshold int       `json:"threshold"`
}

type UserParams struct {
  Username string `json:"username"`
  Password string `json:"password"`
}
type User struct {
  Jsonrpc string      `json:"jsonrpc"`
  Method  string      `json:"method"`
  Params  UserParams `json:"params"`
  ID      int         `json:"id"`
}

type Params struct {
  Assetid   string  `json:"assetID"`
  Payload   string  `json:"payload"`
  To        string  `json:"to"`
  Name      string  `json:"name"`
  Symbol    string  `json:"symbol"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
  From      []string  `json:"from"`
  Changeaddr  string  `json:"changeAddr"`
  Mintersets  []Mintersets  `json:"minterSets"`
}

type URI struct {
	Address string
	Port    string
}

/*
  Utility functions
*/
func SendRequest(uri URI, payloadBytes []byte) {
	node := uri.Address + ":" + uri.Port + "/ext/bc/X"
	body := bytes.NewReader(payloadBytes)
	// req, err := http.NewRequest("POST", node, body)
  req, err := http.NewRequest("POST", node, body)
  if (err != nil ) {
    // Log error
  }

  req.Header.Set("Content-Type", "application/json;")
  resp, err := http.DefaultClient.Do(req)
  if (err != nil ) {
    // Log error
  }
  defer resp.Body.Close()
}

/*
  X-Chain NFT Functions

  Current implemented functions
  - CreateNFTAsset
  - MintNFTAsset
  - SendNFT
*/
func CreateNFTAsset(data Payload, uri URI) {
  // TODO: Check payload params for required params
  // TODO: Check if users has balance to mint assets
  payloadBytes, err := json.Marshal(data)
  if err != nil {
		// Log Error
  }

  SendRequest(uri, payloadBytes)
}

//Mint NFT
func MintNFTAsset(data Payload, uri URI) {
  // TODO: Check payload params for required params
  // TODO: Form execution and call from createNFTAsset

}

func SendNFT(data Payload, uri URI) {
  // TODO: Check payload params for required params
  // TODO: Check if users owns NFT

}

// Get Balance

// Send NFT

/*
  User related functions

  Current implemented functions
  - getAddress
  - CreateUser
  - CreateAddress{Mnemonic,Request}
*/
// Get X-Chain Adress from mnemonic seed (requires cchain import)
// TODO: Pass string as string and convert to []byte
func GetAddressMnemonic(seed []byte) *bip32.Key {
  if (seed != nil) {
    // Refer to 
    // https://github.com/ava-labs/avax-js-cli-tools/blob/3e3f714e4227aca83dc3978fcb6a4fd698e09065/address_gen.js#L60
    if (bip39.IsMnemonicValid(string(seed[:]))) {
      masterKey, _ := bip32.NewMasterKey(seed)
      publicKey := masterKey.PublicKey()
      return publicKey
    }
  }
  return nil
}

// Get address from user password combination
func GetAddressRequest(data Payload, uri URI) string {
  // TODO: Recreate payload and send

  return ""
}

// CreateUser function
func CreateUser(data Payload, uri URI) {
  // Create address using json request
  user := User{
    Jsonrpc: data.Jsonrpc,
    Method: data.Method,
    ID: data.ID,
    Params: UserParams{
      Username: data.Params.Username,
      Password: data.Params.Password,
    },
  }

	payloadBytes, err := json.Marshal(user)
  if ( err != nil ) {
    // log er
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
func CreateAddress(data Payload, uri URI, seed []byte) (*bip32.Key) {
  // TODO: Check for malformed data
  // Derive Address from key
  if (seed != nil ) {
    if (bip39.IsMnemonicValid(string(seed[:]))) {
      masterKey, _ := bip32.NewMasterKey(seed)
      publicKey := masterKey.PublicKey()
      return publicKey
    }
    // TODO: Log error that mnemonic is not valid
  }

  // Create address using json request
  user := User{
    Jsonrpc: data.Jsonrpc,
    Method: data.Method,
    ID: data.ID,
    Params: UserParams{
      Username: data.Params.Username,
      Password: data.Params.Password,
    },
  }

	payloadBytes, err := json.Marshal(user)
  if ( err != nil ) {
    // log error
  }
  SendRequest(uri, payloadBytes)
  // Return key generated from address
  return nil
}
