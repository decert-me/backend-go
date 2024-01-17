package initialize

import (
	"backend-go/internal/job/config"
	"backend-go/pkg/balancer"
)

func InitProviderV2(c *config.Config) map[int64]*balancer.SmoothRoundrobin {
	rpcV2 := make(map[int64]*balancer.SmoothRoundrobin)
	for chainID, contract := range c.ContractV2 {
		w := balancer.NewSmoothRoundrobin()
		for _, provider := range contract.Provider {
			w.Add(provider.Url, provider.Weight)
		}
		rpcV2[chainID] = w
	}
	// version 1
	w := balancer.NewSmoothRoundrobin()
	for _, provider := range c.BlockChain.Provider {
		w.Add(provider.Url, provider.Weight)
	}
	rpcV2[0] = w
	return rpcV2
}
