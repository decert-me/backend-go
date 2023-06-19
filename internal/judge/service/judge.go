package service

import (
	"backend-go/internal/judge/model/request"
	"backend-go/internal/judge/model/response"
	"errors"
)

func (s *Service) TryRun(address string, req request.TryRunReq) (tryRunRes response.TryRunRes, err error) {
	if req.Lang == "Solidity" {
		return s.SolidityTryRun(address, req)
	} else if req.Lang == "JavaScript" {
		return s.JavaScriptTryRun(req)
	} else if req.Lang == "Golang" {
		return s.GolangTryRun(req)
	} else if req.Lang == "Python" {
		return s.PythonTryRun(req)
	} else if req.Lang == "TypeScript" {
		return s.TypeScriptTryRun(req)
	} else if req.Lang == "Move" {
		return s.MoveTryRun(address, req)
	} else {
		return tryRunRes, errors.New("暂不支持的语言")
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
	} else if req.Lang == "Python" {
		return s.PythonTryTestRun(req)
	} else if req.Lang == "TypeScript" {
		return s.TypeScriptTryTestRun(req)
	} else if req.Lang == "Move" {
		return s.MoveTryTestRun(address, req)
	} else {
		return tryRunRes, errors.New("暂不支持的语言")
	}
	return
}
