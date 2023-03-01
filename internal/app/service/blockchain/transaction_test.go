package blockchain

import (
	"backend-go/internal/app/config"
	"backend-go/internal/app/dao"
	"backend-go/internal/app/model"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"sync/atomic"
	"testing"
)

func TestBlockChain_StartTransaction(t *testing.T) {
	type fields struct {
		c             *config.Config
		dao           *dao.Dao
		client        *ethclient.Client
		traversed     atomic.Bool
		TaskChain     chan model.Transaction
		contractEvent map[common.Hash]string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BlockChain{
				c:             tt.fields.c,
				dao:           tt.fields.dao,
				client:        tt.fields.client,
				traversed:     tt.fields.traversed,
				TaskChain:     tt.fields.TaskChain,
				contractEvent: tt.fields.contractEvent,
			}
			b.StartTransaction()
		})
	}
}
