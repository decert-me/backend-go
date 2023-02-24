package blockchain

import (
	ABI "backend-go/abi"
	"backend-go/internal/app/global"
	"backend-go/internal/app/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"log"
	"math/big"
	"testing"
)

func TestAirdropBadge(t *testing.T) {
	global.CONFIG.BlockChain.PrivateKey = "94e0a5961679f979d86020010513d0825d1cb6905e6ae0bf31f41e7fc23dd272"
	global.CONFIG.Contract.Badge = "0x0049770260b599Ecc2e2c0645450c965A44938b7"
	global.CONFIG.Contract.Quest = "0xfE3e0366a52C6F668a1026dAF5e81162d34Ec38b"
	global.CONFIG.Contract.QuestMinter = "0xbE866FE4BAFC11ae886238772AFBD24570f9B530"
	privateKey, err := crypto.HexToECDSA(global.CONFIG.BlockChain.PrivateKey)
	if err != nil {
		return
	}
	address, err := utils.PrivateKeyToAddress(privateKey)
	if err != nil {
		return
	}
	hash := solsha3.SoliditySHA3(
		// types
		[]string{"uint32", "uint32", "uint192", "string", "string", "address", "address"},
		// values
		[]interface{}{
			0, 0, 0, "Airdrop", "ipfs://QmWKAEJq72U57f6qbo3Syg3CYcMMG1LgNfZqxXMZXhVwkW", global.CONFIG.Contract.QuestMinter, address,
		},
	)
	prefixedHash := solsha3.SoliditySHA3WithPrefix(hash)
	signature, err := crypto.Sign(prefixedHash, privateKey)
	signature[64] += 27

	//AirdropBadge()
	client, err := ethclient.Dial("https://rpc.ankr.com/eth_goerli")
	if err != nil {
		panic("ethclient dial error")
	}
	instance, err := ABI.NewQuestMinter(common.HexToAddress(global.CONFIG.Contract.QuestMinter), client)
	if err != nil {
		log.Fatal(err)
	}
	quest, err := ABI.NewQuest(common.HexToAddress(global.CONFIG.Contract.Quest), client)
	if err != nil {
		log.Fatal(err)
	}
	badge, err := ABI.NewBadge(common.HexToAddress(global.CONFIG.Contract.Badge), client)
	if err != nil {
		log.Fatal(err)
	}
	questData := ABI.IQuestQuestData{
		StartTs: 0,
		EndTs:   0,
		Supply:  big.NewInt(0),
		Title:   "Airdrop",
		Uri:     "ipfs://QmWKAEJq72U57f6qbo3Syg3CYcMMG1LgNfZqxXMZXhVwkW",
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(5))
	if err != nil {
		return
	}
	transactOpts := &bind.TransactOpts{
		From:     address,
		Nonce:    auth.Nonce,
		Signer:   auth.Signer,
		Value:    big.NewInt(0),
		GasPrice: auth.GasPrice,
		GasLimit: auth.GasLimit,
		Context:  auth.Context,
		NoSend:   false,
	}
	//res, err := instance.Signer(nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(res)
	_, _, _, _, _, _ = transactOpts, signature, questData, instance, quest, badge
	//fmt.Println(address)
	//res, err := quest.Minters(nil, address)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(res)
	//
	//res2, err := quest.Minters(nil, common.HexToAddress("0xbE866FE4BAFC11ae886238772AFBD24570f9B530"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(res2)
	//res, err := quest.SetMinter(transactOpts, common.HexToAddress("0xbE866FE4BAFC11ae886238772AFBD24570f9B530"), true)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(res.Hash())
	//fmt.Println(hexutil.Encode(signature))
	tx, err := instance.CreateQuest(transactOpts, questData, signature)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(tx)
	fmt.Println(tx.Hash())
	//res, err := badge.Uri(nil, big.NewInt(10003))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(res)
	//res, err := instance.BalanceOf(nil, common.HexToAddress("0x7d32D1DE76acd73d58fc76542212e86ea63817d8"), big.NewInt(10065))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf(res.String())
}
