package utils

import (
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"regexp"
)

// PrivateKeyToAddress
// @description: 私钥转地址
// @param: privateKey *ecdsa.PrivateKey
// @return: common.Address, error
func PrivateKeyToAddress(privateKey *ecdsa.PrivateKey) (common.Address, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.HexToAddress(""), errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	return common.HexToAddress(crypto.PubkeyToAddress(*publicKeyECDSA).Hex()), nil
}

// VerifySignature
// @description: 验证签名消息
// @param: from, sigHex string, msg []byte
// @return: bool
func VerifySignature(from, sigHex string, msg []byte) bool {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	sig := hexutil.MustDecode(sigHex)

	msg = accounts.TextHash(msg)
	if sig[crypto.RecoveryIDOffset] == 27 || sig[crypto.RecoveryIDOffset] == 28 {
		sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	recovered, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return false
	}
	recoveredAddr := crypto.PubkeyToAddress(*recovered)

	return from == recoveredAddr.Hex()
}

// IsValidAddress validate hex address
func IsValidAddress(iaddress interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := iaddress.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}
