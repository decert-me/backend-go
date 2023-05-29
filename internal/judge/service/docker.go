package service

import (
	"backend-go/pkg/log"
	"go.uber.org/zap"
	"strings"
)

// ExistsDocker 检测 docker 是否存在
func ExistsDocker(name string) bool {
	args := []string{"container", "inspect", name}
	execRes, err := execCommand("", "docker", args...)
	if err != nil {
		log.Errorv("execCommand error", zap.Error(err))
		return false
	}
	if strings.Contains(execRes, "No such container") {
		return false
	}
	return true
}

// CreateDocker 创建docker
func CreateDocker(name string, image string, argList []string) bool {
	args := []string{"run", "-itd", "--name", name, image, "/bin/bash"}
	args = append(args, argList...)
	execRes, err := execCommand("", "docker", args...)
	if err != nil {
		log.Errorv("execCommand error", zap.Error(err))
		return false
	}
	if strings.Contains(execRes, "No such container") {
		return false
	}
	return true
}

// DelDocker 删除容器
func DelDocker(name string) bool {
	_, err := execCommand("", "docker", "stop", name)
	if err != nil {
		log.Errorv("execCommand error", zap.Error(err))
		return false
	}
	execRm, err := execCommand("", "docker", "rm", name)
	if err != nil {
		log.Errorv("execCommand error", zap.Error(err))
		return false
	}
	if strings.Contains(execRm, "No such container") {
		return false
	}
	return true
}
