package service

import "backend-go/internal/app/model/response"

// GetAddressChallengeCount 获取地址完成挑战/获得NFT的数量
func (s *Service) GetAddressChallengeCount(address string) (res response.GetAddressChallengeCountRes, err error) {
	return s.dao.GetAddressChallengeCount(address)
}
