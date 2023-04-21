package service

import (
	"backend-go/internal/app/model/response"
	"context"
	"encoding/json"
	"errors"
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
			return response.GetEnsResponse{}, nil
		}
		_ = json.Unmarshal([]byte(ensValue), &result)
		return result, err
	}
	if strings.Contains(q, ".") {
		return s.resolveName(ctx, q)
	} else {
		if !common.IsHexAddress(q) {
			return result, errors.New("RecordNotFound")
		}
		return s.resolveAddress(ctx, common.HexToAddress(q).Hex())
	}
}

func (s *Service) resolveName(ctx context.Context, input string) (result response.GetEnsResponse, err error) {
	result.Domain = input
	client, err := ethclient.Dial(s.c.BlockChain.EnsRpc)
	if err != nil {
		return result, errors.New("UnexpectedError")
	}
	// Resolve a name to an address.
	address, err := ens.Resolve(client, input)
	if err != nil {
		jsonStr, err := json.Marshal(result)
		if err != nil {
			return result, errors.New("UnexpectedError")
		}
		s.dao.SetEns(ctx, input, string(jsonStr), time.Second*60)
		return result, nil
	}
	result.Address = address.Hex()
	text, errResolver := ens.NewResolver(client, input)
	if errResolver == nil {
		avatar, errAvatar := text.Text("avatar")
		if errAvatar == nil {
			result.Avatar = avatar
		}
	}

	jsonStr, err := json.Marshal(result)
	if err != nil {
		return result, errors.New("UnexpectedError")
	}
	if err = s.dao.SetEns(ctx, input, string(jsonStr), time.Hour*24); err != nil {
		return result, errors.New("UnexpectedError")
	}
	return result, nil
}

func (s *Service) resolveAddress(ctx context.Context, input string) (result response.GetEnsResponse, err error) {
	result.Address = input
	client, err := ethclient.Dial(s.c.BlockChain.EnsRpc)
	if err != nil {
		return result, errors.New("UnexpectedError")
	}
	// Resolve address to name
	domain, err := ens.ReverseResolve(client, common.HexToAddress(input))
	if err != nil {
		jsonStr, err := json.Marshal(result)
		if err != nil {
			return result, errors.New("UnexpectedError")
		}
		s.dao.SetEns(ctx, input, string(jsonStr), time.Second*60)
		return result, nil
	}
	result.Domain = domain
	text, errResolver := ens.NewResolver(client, domain)
	if errResolver == nil {
		avatar, errAvatar := text.Text("avatar")
		if errAvatar == nil {
			result.Avatar = avatar
		}
	}

	jsonStr, err := json.Marshal(result)
	if err != nil {
		return result, errors.New("UnexpectedError")
	}
	if err = s.dao.SetEns(ctx, input, string(jsonStr), time.Hour*24); err != nil {
		return result, errors.New("UnexpectedError")
	}
	return result, nil
}
