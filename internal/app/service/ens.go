package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/response"
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	ens "github.com/wealdtech/go-ens/v3"
	"strings"
	"time"
)

func (s *Service) GetEnsRecords(ctx context.Context, q string) (result interface{}, err error) {
	// Get records from Cache
	if strings.Contains(q, ".") {
		res, errNil := s.dao.GetAddressByEns(q)
		if errNil != nil {
			return s.resolveName(ctx, q)
		}
		if res.UpdatedAt.Before(time.Now().Add(-1)) {
			go s.resolveName(ctx, q)
		}
		return response.GetEnsResponse{Address: res.Address, Domain: res.Domain, Avatar: res.Avatar}, nil
	} else {
		if !common.IsHexAddress(q) {
			return result, errors.New("RecordNotFound")
		}
		res, errNil := s.dao.GetEnsByAddress(q)
		if errNil != nil {
			return s.resolveAddress(ctx, common.HexToAddress(q).Hex())
		}
		if res.UpdatedAt.Before(time.Now().Add(-1)) {
			go s.resolveAddress(ctx, common.HexToAddress(q).Hex())
		}
		return response.GetEnsResponse{Address: res.Address, Domain: res.Domain, Avatar: res.Avatar}, nil
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
	if err != nil {
		return result, errors.New("UnexpectedError")
	}
	if err = s.dao.SetEns(model.Ens{Address: address.Hex(), Domain: input, Avatar: result.Avatar}); err != nil {
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

	if err != nil {
		return result, errors.New("UnexpectedError")
	}
	if err = s.dao.SetEns(model.Ens{Address: input, Domain: domain, Avatar: result.Avatar}); err != nil {
		return result, errors.New("UnexpectedError")
	}
	return result, nil
}
