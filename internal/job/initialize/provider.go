package initialize

import (
	"backend-go/internal/job/config"
)

func InitProvider(c *config.Config) map[int]string {
	temp := make(map[int]string)
	for _, v := range c.BlockChain.Provider {
		temp[v.ChainID] = v.Url[0]
	}
	return temp
}
