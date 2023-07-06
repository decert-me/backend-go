package initialize

import (
	"backend-go/internal/app/config"
	"backend-go/pkg/balancer"
)

func InitJudge(c *config.Config) *balancer.SmoothRoundrobin {
	w := balancer.NewSmoothRoundrobin()
	for _, api := range c.Judge.JudgeApi {
		w.Add(api.Url, api.Weight)
	}
	return w
}
