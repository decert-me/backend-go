package service

import (
	"backend-go/pkg/log"
	"errors"
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
func CreateDocker(name string, image string, argList []string) error {
	args := []string{"run", "-itd", "--name", name}
	args = append(args, argList...)
	args = append(args, image)
	execRes, err := execCommand("", "docker", args...)
	if err != nil {
		log.Errorv("execCommand error", zap.Error(err))
		return errors.New("execCommand error")
	}
	if strings.Contains(execRes, "Unable to find image") {
		log.Errorv("image not found", zap.String("image", image))
		return errors.New("image not found")
	}
	return nil
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
