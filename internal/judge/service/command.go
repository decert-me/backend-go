package service

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

func execCommand(dir string, command string, args ...string) (res string, err error) {
	cmd := exec.Command(command, args...)
	fmt.Println(cmd.Args) //显示运行的命令
	cmd.Dir = dir
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return res, err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return res, err
	}
	err = cmd.Start()
	if err != nil {
		return res, err
	}
	oReader := bufio.NewReader(stdout)
	eReader := bufio.NewReader(stderr)
	var stringBuf strings.Builder
	for {
		line, err := oReader.ReadString('\n')
		if err != nil {
			break
		}
		stringBuf.WriteString(line)
	}

	for {
		line, err := eReader.ReadString('\n')
		if err != nil {
			break
		}
		stringBuf.WriteString(line)
	}

	return stringBuf.String(), nil
}
