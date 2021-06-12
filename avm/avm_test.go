package goavx

import (
	"reflect"
	"testing"

	. "github.com/smartpassnft/goavx/avm"
	. "github.com/smartpassnft/goavx/avm/utils"
	"github.com/tyler-smith/go-bip32"
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
		data Payload
		uri  URI
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
		data Payload
		uri  URI
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
		data Payload
		uri  URI
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
		data    Payload
		address []string
		uri     URI
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetUTXOS(tt.args.data, tt.args.address, tt.args.uri)
		})
	}
}

func TestCreateNFTAsset(t *testing.T) {
	type args struct {
		data Payload
		uri  URI
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateNFTAsset(tt.args.data, tt.args.uri)
		})
	}
}

func TestSendRequest(t *testing.T) {
	type args struct {
		uri          URI
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
			SendRequest(tt.args.uri, tt.args.payloadBytes)
		})
	}
}
