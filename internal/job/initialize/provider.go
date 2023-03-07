package initialize

import (
	"backend-go/internal/job/config"
	"backend-go/pkg/balancer"
)

func InitProvider(c *config.Config) *balancer.SmoothRoundrobin {
	w := balancer.NewSmoothRoundrobin()
	for _, provider := range c.BlockChain.Provider {
		w.Add(provider.Url, provider.Weight)
	}
	return w
}
