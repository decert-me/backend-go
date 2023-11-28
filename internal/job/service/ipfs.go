package service

import (
	"fmt"
	"github.com/imroc/req/v3"
	"time"
)

// GetDataFromCid
// @description: 获取IPFS内容
// @param: cid string
// @return: string, error
func (s *Service) GetDataFromCid(cid string) (string, error) {
	baseURL := s.c.IPFS.API
	url := baseURL + cid
	client := req.C().SetTimeout(180 * time.Second).SetCommonRetryCount(1)
	fmt.Println("url", url)
	req, err := client.R().Get(url)
	return req.String(), err
}
