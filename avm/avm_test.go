package avm

import (
	"reflect"
	"testing"

	"github.com/smartpassnft/goavx/avm/utils"
	bip32 "github.com/tyler-smith/go-bip32"
)

func TestCreateAddressSeed(t *testing.T) {
	type args struct {
		seed []byte
	}
	tests := []struct {
		name string
		args args
		want *bip32.Key
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateAddressSeed(tt.args.seed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAddressSeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateAddressAccount(t *testing.T) {
	type args struct {
		data utils.Payload
		uri  utils.URI
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateAddressAccount(tt.args.data, tt.args.uri)
		})
	}
}

func TestGenerateAddress(t *testing.T) {
	type args struct {
		pass string
		num  int
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 *bip32.Key
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GenerateAddress(tt.args.pass, tt.args.num)
			if got != tt.want {
				t.Errorf("GenerateAddress() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GenerateAddress() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	type args struct {
		data utils.Payload
		uri  utils.URI
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateUser(tt.args.data, tt.args.uri)
		})
	}
}

func TestGetAddress(t *testing.T) {
	type args struct {
		data utils.Payload
		uri  utils.URI
		seed []byte
	}
	tests := []struct {
		name string
		args args
		want *bip32.Key
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAddress(tt.args.data, tt.args.uri, tt.args.seed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUTXOS(t *testing.T) {
	type args struct {
		data    utils.Payload
		address []string
		uri     utils.URI
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utils.GetUTXOS(tt.args.data, tt.args.address, tt.args.uri)
		})
	}
}

func TestCreateNFTAsset(t *testing.T) {
	type args struct {
		data utils.Payload
		uri  utils.URI
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utils.CreateNFTAsset(tt.args.data, tt.args.uri)
		})
	}
}

func TestSendRequest(t *testing.T) {
	type args struct {
		uri          utils.URI
		payloadBytes []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utils.SendRequest(tt.args.uri, tt.args.payloadBytes)
		})
	}
}
