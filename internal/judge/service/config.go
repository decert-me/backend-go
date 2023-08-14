package service

import (
	"backend-go/internal/judge/model/judge"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"regexp"
	"strings"
)

type JavaScriptRunCopyIn struct {
	RunMainCode judge.RunMainCode `json:"main.js"`
}

type TypeScriptRunCopyIn struct {
	RunMainCode judge.RunMainCode `json:"main.ts"`
}

type GolangRunCopyIn struct {
	RunMainCode judge.RunMainCode `json:"main.go"`
}

type PythonRunCopyIn struct {
	RunMainCode judge.RunMainCode `json:"main.py"`
}

func (s *Service) JavaScriptRun(code string, codeSnippet string, inputs []string) (result interface{}, err error) {
	// 获取函数名称
	var functionName string
	re := regexp.MustCompile(`function\s+(\w+)\s*\(`)
	matches := re.FindStringSubmatch(codeSnippet)
	if len(matches) > 1 {
		functionName = matches[1]
	} else {
		return result, errors.New("函数名称获取失败")
	}

	var cmdList []judge.RunCmd
	for _, input := range inputs {
		// 输出函数
		inputCode := fmt.Sprintf("\nconsole.log(%s(%s));", functionName, input)
		code = code + inputCode
		// 输入
		runFiles := []judge.RunFiles{
			{Content: input},
			{Name: "stdout", Max: 10240},
			{Name: "stderr", Max: 10240},
		}
		// 代码
		runCopyIn := JavaScriptRunCopyIn{
			RunMainCode: judge.RunMainCode{Content: code},
		}
		//
		cmd := judge.RunCmd{
			Args:        []string{s.c.Judge.JavaScriptPath, "main.js"},
			Env:         []string{"PATH=/usr/bin:/bin"},
			Files:       runFiles,
			CPULimit:    3000000000,
			ClockLimit:  4000000000,
			MemoryLimit: 104857600,
			ProcLimit:   50,
			CPURate:     0.1,
			CopyIn:      runCopyIn,
		}
		cmdList = append(cmdList, cmd)
	}

	return judge.Run{Cmd: cmdList}, nil
}

func (s *Service) TypeScriptRun(code string, codeSnippet string, inputs []string) (result interface{}, err error) {
	// 获取函数名称
	var functionName string
	re := regexp.MustCompile(`function\s+(\w+)\s*\(`)
	matches := re.FindStringSubmatch(codeSnippet)
	if len(matches) > 1 {
		functionName = matches[1]
	} else {
		return result, errors.New("函数名称获取失败")
	}

	var cmdList []judge.RunCmd
	for _, input := range inputs {
		// 输出函数
		inputCode := fmt.Sprintf("\nconsole.log(%s(%s));", functionName, input)
		code = code + inputCode
		// 输入
		runFiles := []judge.RunFiles{
			{Content: input},
			{Name: "stdout", Max: 10240},
			{Name: "stderr", Max: 10240},
		}
		// 代码
		runCopyIn := TypeScriptRunCopyIn{
			RunMainCode: judge.RunMainCode{Content: code},
		}
		fmt.Println()
		cmd := judge.RunCmd{
			Args:        []string{s.c.Judge.TypeScriptPath, "main.ts"},
			Env:         []string{"PATH=/usr/bin:/bin:/usr/local/bin/"},
			Files:       runFiles,
			CPULimit:    3000000000,
			ClockLimit:  4000000000,
			MemoryLimit: 1048576000,
			ProcLimit:   50,
			CPURate:     0.1,
			CopyIn:      runCopyIn,
		}
		cmdList = append(cmdList, cmd)
	}

	return judge.Run{Cmd: cmdList}, nil
}

func (s *Service) GolangRun(code string, codeSnippet string, inputs []string) (result interface{}, err error) {
	// 获取函数名称
	var functionName string
	var parameter string
	re := regexp.MustCompile(`func\s+(\w+)\((.*?)\)`)
	matches := re.FindStringSubmatch(codeSnippet)
	if len(matches) > 2 {
		functionName = matches[1]
		parameter = matches[2]
	} else {
		return result, errors.New("函数名称获取失败")
	}
	// 获取变量类型
	var varTypes []string
	//fmt.Println(strings.Split(parameter, ","))
	var markIndex = -1
	for i, v := range strings.Split(parameter, ",") {
		v = strings.TrimSpace(v)
		if len(strings.Split(v, " ")) == 1 {
			if markIndex == -1 {
				markIndex = i
			}
			continue
		}
		varType := strings.Split(v, " ")[1]
		// func twoSum(nums , target int) 情况
		if markIndex != -1 {
			for start := markIndex; start <= i; start++ {
				varTypes = append(varTypes, varType)
			}
			markIndex = -1
		} else {
			varTypes = append(varTypes, varType)
		}
	}
	var cmdList []judge.RunCmd
	for _, input := range inputs {
		// 输出函数
		inputStr := strings.Builder{}
		for i, v := range gjson.Parse("[" + input + "]").Array() {
			if len(v.String()) != 0 && v.String()[0] == '[' {
				inputStr.WriteString(strings.Replace(strings.Replace(v.String(), "]", "}", 1), "[", fmt.Sprintf("%s{", varTypes[i]), 1))
			} else {
				inputStr.WriteString(v.String())
			}
			if i+1 < len(gjson.Parse("["+input+"]").Array()) {
				inputStr.WriteString(",")
			}
		}
		fmt.Println("inputStr", inputStr.String())
		//input := strings.Replace(input, "[", "")
		inputCode := fmt.Sprintf("package main\nimport \"fmt\"\nfunc main() {\n\tfmt.Println(%s(%s))\n}\n%s", functionName, inputStr.String(), code)
		// 输入
		runFiles := []judge.RunFiles{
			{Content: input},
			{Name: "stdout", Max: 10240},
			{Name: "stderr", Max: 10240},
		}
		// 代码
		runCopyIn := GolangRunCopyIn{
			RunMainCode: judge.RunMainCode{Content: inputCode},
		}
		cmd := judge.RunCmd{
			Args:        []string{s.c.Judge.GolangPath, "run", "main.go"},
			Env:         []string{"PATH=/usr/bin:/bin", "GOCACHE=/tmp"},
			Files:       runFiles,
			CPULimit:    3000000000,
			ClockLimit:  4000000000,
			MemoryLimit: 104857600,
			ProcLimit:   100,
			CPURate:     0.1,
			CopyIn:      runCopyIn,
		}
		cmdList = append(cmdList, cmd)
	}

	return judge.Run{Cmd: cmdList}, nil
}

func (s *Service) PythonScriptRun(code string, codeSnippet string, inputs []string) (result interface{}, err error) {
	// 获取函数名称
	var functionName string
	re := regexp.MustCompile(`def\s+(\w+)\s*\(`)
	matches := re.FindStringSubmatch(codeSnippet)
	if len(matches) > 1 {
		functionName = matches[1]
	} else {
		return result, errors.New("函数名称获取失败")
	}

	var cmdList []judge.RunCmd
	for _, input := range inputs {
		// 输出函数
		inputCode := fmt.Sprintf("\nprint(%s(%s))", functionName, input)
		code = code + inputCode
		// 输入
		runFiles := []judge.RunFiles{
			{Content: input},
			{Name: "stdout", Max: 10240},
			{Name: "stderr", Max: 10240},
		}
		// 代码
		runCopyIn := PythonRunCopyIn{
			RunMainCode: judge.RunMainCode{Content: code},
		}
		//
		fmt.Println()
		cmd := judge.RunCmd{
			Args:        []string{s.c.Judge.PythonPath, "main.py"},
			Env:         []string{"PATH=/usr/bin:/bin"},
			Files:       runFiles,
			CPULimit:    3000000000,
			ClockLimit:  4000000000,
			MemoryLimit: 104857600,
			ProcLimit:   50,
			CPURate:     0.1,
			CopyIn:      runCopyIn,
		}
		cmdList = append(cmdList, cmd)
	}

	return judge.Run{Cmd: cmdList}, nil
}
