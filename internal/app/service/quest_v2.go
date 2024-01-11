package service

import (
	"backend-go/internal/app/model/request"
	"errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/liangjies/go-solidity-sha3"
	"math/big"
)

func (s *Service) AddQuestV2(address string, add request.AddQuestV2Request) (res string, err error) {
	privateKey, err := crypto.HexToECDSA(s.c.BlockChain.SignPrivateKey)
	if err != nil {
		return
	}
	hash := solsha3.SoliditySHA3(
		// types
		[]string{"uint256", "uint32", "uint32", "string", "string", "address", "address"},
		// values
		[]interface{}{
			add.ChainID, add.StartTs, add.EndTs, add.Title, add.Uri, s.c.ContractV2[add.ChainID].QuestMinter, address,
		},
	)
	prefixedHash := solsha3.SoliditySHA3WithPrefix(hash)
	signature, err := crypto.Sign(prefixedHash, privateKey)
	signature[64] += 27
	return hexutil.Encode(signature), err
}

func (s *Service) UpdateQuestV2(address string, modify request.UpdateQuestV2Request) (res string, err error) {
	privateKey, err := crypto.HexToECDSA(s.c.BlockChain.SignPrivateKey)
	if err != nil {
		return
	}
	tokenId, set := big.NewInt(0).SetString(modify.TokenId, 10)
	if !set {
		return res, errors.New("TokenIDInvalid")
	}
	hash := solsha3.SoliditySHA3(
		// types
		[]string{"uint256", "uint32", "uint32", "string", "string", "address", "address"},
		// values
		[]interface{}{
			tokenId, modify.StartTs, modify.EndTs, modify.Title, modify.Uri, s.c.ContractV2[modify.ChainID].QuestMinter, address,
		},
	)
	prefixedHash := solsha3.SoliditySHA3WithPrefix(hash)
	signature, err := crypto.Sign(prefixedHash, privateKey)
	signature[64] += 27
	return hexutil.Encode(signature), err
}
