package utils

import (
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gagliardetto/solana-go"
	"regexp"
)

// PrivateKeyToAddress
// @description: 私钥转地址
// @param: privateKey *ecdsa.PrivateKey
// @return: common.Address, error
func PrivateKeyToAddress(privateKey *ecdsa.PrivateKey) (address common.Address, err error) {
	defer func() {
		if errRv := recover(); errRv != nil {
			err = errors.New("error")
			return
		}
	}()
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

// VerifySignatureSolana
// @description: 验证Solana签名消息
// @param: from, sigHex string, msg []byte
// @return: bool
func VerifySignatureSolana(from, sigHex string, message []byte) bool {
	signature, err := solana.SignatureFromBase58(sigHex)
	if err != nil {
		return false
	}
	pubKey, err := solana.PublicKeyFromBase58(from)
	if err != nil {
		return false
	}

	if !signature.Verify(pubKey, message) {
		return false
	}

	return true
}

// IsValidAddress validate hex address
func IsValidAddress(address interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := address.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}
