package service

import (
	"backend-go/internal/judge/model/request"
	"backend-go/internal/judge/model/response"
	"github.com/tidwall/gjson"
	"strings"
)

func (s *Service) CastCall(req request.CastCallReq) (res response.CastCallRes, err error) {
	args := []string{"call", req.To}
	if req.CallData != "" {
		args = append(args, req.CallData)
	} else if !strings.Contains(req.Data, "],") {
		args = append(args, req.Method)
		args = append(args, strings.Split(req.Data, ",")...)
	} else {
		args = append(args, req.Method)
		for _, v := range strings.Split(req.Data, "],") {
			if v[:1] == "[" {
				args = append(args, v+"]")
				continue
			} else {
				args = append(args, strings.Split(v, ",")...)
			}
		}
	}
	execRes, err := execCommand("", "cast", args...)
	if err != nil {
		res.Msg = err.Error()
		res.Status = 1
		return
	}
	if len(execRes) > 5 && execRes[:5] == "Error" {
		res.Msg = execRes
		res.Status = 1
		return
	}
	res.Data = strings.TrimRight(execRes, "\n")
	return res, nil
}

func (s *Service) CastSend(req request.CastSendReq) (res response.CastSend, err error) {
	args := []string{"send", req.To}
	if req.CallData != "" {
		args = append(args, req.CallData)
	} else {
		args = append(args, req.Method, req.Data)
	}
	args = append(args, "--private-key=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80", "--json")
	execRes, err := execCommand("", "cast", args...)
	if err != nil {
		return
	}
	if !gjson.Valid(execRes) {
		res.Msg = execRes
		return
	}
	res.GasUsed = gjson.Get(execRes, "gasUsed").String()
	res.Status = gjson.Get(execRes, "status").String()
	return res, nil

}
