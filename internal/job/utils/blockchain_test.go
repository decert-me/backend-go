package utils

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"reflect"
	"testing"
)

func TestIsValidAddress(t *testing.T) {
	type args struct {
		address interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "#0 empty address should false", args: args{address: ""}, want: false},
		{name: "#1 address length not enough should false", args: args{address: "0x7d32D1DE76acd73d58fc76542212e86ea63817d"}, want: false},
		{name: "#2 zero address should pass", args: args{address: "0x0000000000000000000000000000000000000000"}, want: true},
		{name: "#3 valid address should pass", args: args{address: "0x7d32D1DE76acd73d58fc76542212e86ea63817d8"}, want: true},
		{name: "#4 common.Address should pass", args: args{address: common.HexToAddress("0x7d32D1DE76acd73d58fc76542212e86ea63817d8")}, want: true},
		{name: "#5 invalid Address should fail", args: args{address: 1243324324}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidAddress(tt.args.address); got != tt.want {
				t.Errorf("IsValidAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrivateKeyToAddress(t *testing.T) {
	// init private key
	// privateKey for testing don't use in production
	privateKey, _ := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	privateKeyInvalid, _ := crypto.HexToECDSA("fad9c840a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a29")
	type args struct {
		privateKey *ecdsa.PrivateKey
	}
	tests := []struct {
		name    string
		args    args
		want    common.Address
		wantErr bool
	}{
		{name: "#0 privateKey to address should pass", args: args{privateKey: privateKey}, want: common.HexToAddress("0x96216849c49358B10257cb55b28eA603c874b05E"), wantErr: false},
		{name: "#1 invalid privateKey to address should fail", args: args{privateKey: privateKeyInvalid}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrivateKeyToAddress(tt.args.privateKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrivateKeyToAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrivateKeyToAddress() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerifySignature(t *testing.T) {
	type args struct {
		from   string
		sigHex string
		msg    []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "#0 should pass", args: args{"0x0EaE3eF6CC7176553E6B45d94e9eFDE2Da7B82a5", "0x34850b7e36e635783df0563c7202c3ac776df59db5015d2b6f0add33955bb5c43ce35efb5ce695a243bc4c5dc4298db40cd765f3ea5612d2d57da1e4933b2f201b", []byte("Example `personal_sign` message")}, want: true},
		{name: "#1 wrong address should fail", args: args{"0x0EaE3eF6CC7176553E6B45d94e9eFDE2Da7B82a6", "0x34850b7e36e635783df0563c7202c3ac776df59db5015d2b6f0add33955bb5c43ce35efb5ce695a243bc4c5dc4298db40cd765f3ea5612d2d57da1e4933b2f201b", []byte("Example `personal_sign` message")}, want: false},
		{name: "#2 wrong address should fail", args: args{"0x0EaE3eF6CC7176553E6B45d94e9eFDE2Da7B82a6", "2342342432", []byte("Example `personal_sign` message")}, want: false},
		{name: "#3 wrong signature should fail", args: args{"0x0EaE3eF6CC7176553E6B45d94e9eFDE2Da7B82a5", "0x34850b7e36e635783df0563c7202c3ac776df59db5015d2b6f0add33955bb5c43ce35efb5ce695a243bc4c5dc4298db40cd765f3ea5612d2d57da1e4933b2f201c", []byte("Example `personal_sign` message")}, want: false},
		{name: "#4 wrong signature should fail", args: args{"0x0EaE3eF6CC7176553E6B45d94e9eFDE2Da7B82a5", "1234123124124t34t", []byte("Example `personal_sign` message")}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerifySignature(tt.args.from, tt.args.sigHex, tt.args.msg); got != tt.want {
				t.Errorf("VerifySignature() = %v, want %v", got, tt.want)
			}
		})
	}
}
