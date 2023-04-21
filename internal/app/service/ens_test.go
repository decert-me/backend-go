package service

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	ens "github.com/wealdtech/go-ens/v3"
)

func TestEns(t *testing.T) {
	// Replace SECRET with your own access token for this example to work.
	client, err := ethclient.Dial("https://rpc.ankr.com/eth")
	if err != nil {
		panic(err)
	}

	// Resolve a name to an address.
	domain := "brantly.eth"
	address, err := ens.Resolve(client, domain)
	if err != nil {
		panic(err)

	}
	fmt.Printf("Address of %s is %s\n", domain, address.Hex())

	address = common.HexToAddress("0x7d32D1DE76acd73d58fc76542212e86ea63817d8")
	// Reverse resolve an address to a name.
	reverse, err := ens.ReverseResolve(client, address)
	if err != nil {
		panic(err)
	}
	fmt.Println(reverse)
	if reverse == "" {
		fmt.Printf("%s has no reverse lookup\n", address.Hex())
	} else {
		fmt.Printf("Name of %s is %s\n", address.Hex(), reverse)
	}
	dns, err := ens.NewResolver(client, domain)
	if err != nil {
		panic(err)
	}
	fmt.Println(dns.Text("avatar"))
}

func TestService_GetEnsRecords(t *testing.T) {
	fmt.Println(common.IsHexAddress("b3f204a5f3dabef6be51015fd57e307080db6498"))
}
