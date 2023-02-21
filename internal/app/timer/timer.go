package timer

import (
	"backend-go/internal/app/blockchain"
	"github.com/robfig/cron/v3"
)

func Timer() {
	pendingAirdropBadge()
}

func pendingAirdropBadge() {
	t := cron.New()
	t.AddFunc("15 */1 * * *", func() {
		blockchain.AirdropBadge()
	})
	t.Start()
}
