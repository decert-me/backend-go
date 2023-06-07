package service

import (
	"backend-go/internal/judge/model/request"
	"backend-go/internal/judge/model/response"
)

func (s *Service) TryRun(address string, req request.TryRunReq) (tryRunRes response.TryRunRes, err error) {
	if req.Lang == "Solidity" {
		return s.SolidityTryRun(address, req)
	} else if req.Lang == "JavaScript" {
		return s.JavaScriptTryRun(req)
	} else if req.Lang == "Golang" {
		return s.GolangTryRun(req)
	}
	return
}

func (s *Service) TryTestRun(address string, req request.TryTestRunReq) (tryRunRes response.TryRunRes, err error) {
	if req.Lang == "Solidity" {
		return s.SolidityTryTestRun(address, req)
	} else if req.Lang == "JavaScript" {
		return s.JavaScriptTryTestRun(req)
	} else if req.Lang == "Golang" {
		return s.GolangTryTestRun(req)
	}
	return
}
