package service

import (
	"backend-go/internal/judge/model/request"
	"backend-go/internal/judge/model/response"
	"fmt"
	"github.com/tidwall/gjson"
	"strings"
)

func (s *Service) CastCall(ctrName string, req request.CastCallReq) (res response.CastCallRes, err error) {
	args := []string{"call", req.To}
	if req.CallData != "" {
		args = append(args, req.CallData)
	} else {
		args = append(args, "\""+req.Method+"\"")
		inputStr := strings.Builder{}
		for i, v := range gjson.Parse("[" + req.Data + "]").Array() {
			inputStr.WriteString(v.String())
			if i+1 < len(gjson.Parse("["+req.Data+"]").Array()) {
				inputStr.WriteString(" ")
			}
		}
		args = append(args, inputStr.String())
	}
	command := fmt.Sprintf("cast %s", strings.Join(args, " "))
	fmt.Println("command", command)
	argsExec := []string{"exec", "-i", ctrName, "bash", "-c", command}

	execRes, err := execCommand("", "docker", argsExec...)
	fmt.Println("execRes", execRes)
	//execRes, err := execCommand("", "cast", args...)
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

func (s *Service) CastSend(ctrName string, req request.CastSendReq) (res response.CastSend, err error) {
	args := []string{"send", req.To}
	if req.CallData != "" {
		args = append(args, req.CallData)
	} else {
		args = append(args, req.Method, req.Data)
	}
	//args = append(args, "--private-key=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80", "--json")
	privateKey := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	command := fmt.Sprintf("cast %s --private-key=%s --json", strings.Join(args, " "), privateKey)
	argsExec := []string{"exec", "-i", ctrName, "bash", "-c", command}

	execRes, err := execCommand("", "docker", argsExec...)
	//execRes, err := execCommand("", "cast", args...)
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
