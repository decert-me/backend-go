package service

import (
	"backend-go/internal/app/model/request"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/liangjies/go-solidity-sha3"
	"math/big"
)

func (s *Service) UpdateBadgeURIV2(address string, badgeURI request.UpdateBadgeURIRequest) (res string, err error) {
	privateKey, err := crypto.HexToECDSA(s.c.BlockChain.SignPrivateKey)
	if err != nil {
		return
	}
	hash := solsha3.SoliditySHA3(
		// types
		[]string{"uint256", "string", "address", "address"},
		// values
		[]interface{}{
			big.NewInt(badgeURI.TokenId), badgeURI.Uri, s.c.ContractV2[badgeURI.ChainID].Badge, address,
		},
	)
	prefixedHash := solsha3.SoliditySHA3WithPrefix(hash)
	signature, err := crypto.Sign(prefixedHash, privateKey)
	signature[64] += 27
	return hexutil.Encode(signature), err
}
