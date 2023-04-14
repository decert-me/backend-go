package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	ens "github.com/wealdtech/go-ens/v3"
	"strings"
	"time"
)

func (s *Service) GetEnsRecords(ctx context.Context, q string) (result interface{}, err error) {
	// Get records from Cache
	ensValue, err := s.dao.GetEns(ctx, q)
	if err == nil {
		if ensValue == "" {
			return result, errors.New("RecordNotFound")
		}
		_ = json.Unmarshal([]byte(ensValue), &result)
		return result, err
	}
	if strings.Contains(q, ".") {
		return s.resolveName(ctx, q)
	} else {
		return s.resolveAddress(ctx, q)
	}
}

func (s *Service) resolveName(ctx context.Context, input string) (result interface{}, err error) {
	jsonMap := make(map[string]interface{})
	client, err := ethclient.Dial(s.c.BlockChain.EnsRpc)
	if err != nil {
		return result, errors.New("UnexpectedError")
	}
	// Resolve a name to an address.
	address, err := ens.Resolve(client, input)
	if err != nil {
		s.dao.SetEns(ctx, input, "", time.Second*60)
		return result, errors.New("RecordNotFound")
	}
	jsonMap["address"] = address.Hex()
	jsonMap["domain"] = input
	text, errResolver := ens.NewResolver(client, input)
	if errResolver == nil {
		avatar, errAvatar := text.Text("avatar")
		if errAvatar == nil {
			jsonMap["avatar"] = avatar
		}
	}

	jsonStr, err := json.Marshal(jsonMap)
	if err != nil {
		return result, errors.New("UnexpectedError")
	}
	if err = s.dao.SetEns(ctx, input, string(jsonStr), time.Hour*24); err != nil {
		return result, errors.New("UnexpectedError")
	}
	return jsonMap, nil
}

func (s *Service) resolveAddress(ctx context.Context, input string) (result interface{}, err error) {
	jsonMap := make(map[string]interface{})
	client, err := ethclient.Dial(s.c.BlockChain.EnsRpc)
	if err != nil {
		return result, errors.New("UnexpectedError")
	}
	// Resolve address to name
	domain, err := ens.ReverseResolve(client, common.HexToAddress(input))
	if err != nil {
		s.dao.SetEns(ctx, input, "", time.Second*60)
		return result, errors.New("RecordNotFound")
	}
	fmt.Println(domain)
	jsonMap["address"] = input
	jsonMap["domain"] = domain
	text, errResolver := ens.NewResolver(client, domain)
	if errResolver == nil {
		avatar, errAvatar := text.Text("avatar")
		if errAvatar == nil {
			jsonMap["avatar"] = avatar
		}
	}

	jsonStr, err := json.Marshal(jsonMap)
	if err != nil {
		return result, errors.New("UnexpectedError")
	}
	if err = s.dao.SetEns(ctx, input, string(jsonStr), time.Hour*24); err != nil {
		return result, errors.New("UnexpectedError")
	}
	return jsonMap, nil
}
