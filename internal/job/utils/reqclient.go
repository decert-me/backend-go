package utils

import (
	"github.com/imroc/req/v3"
	"sync"
	"time"
)

var clientReq *req.Client
var once sync.Once

func GetReqClient() *req.Client {
	once.Do(func() {
		clientReq = req.C().SetTimeout(30 * time.Second).SetCommonRetryCount(2)
	})
	return clientReq
}
