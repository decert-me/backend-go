package service

import (
	"backend-go/internal/judge/initialize"
	"backend-go/pkg/log"
	"fmt"
	"path"
	"testing"
)

func TestService_ExistsDocker(t *testing.T) {
	// 初始化Viper
	c := initialize.Viper("../cmd/config.yaml")
	// 初始化日志框架
	log.Init(c.Log)
	s := &Service{
		c: c,
	}
	_ = s
	//command := "cd /foundry && forge create " + contract + fmt.Sprintf(" --private-key=%s", private) + " --json"
	//args := []string{"exec", "-i", ctrName, "bash", "-c", command}
	//
	//execRes, err := execCommand("", "docker", args...)
	//CreateDocker("hello", "foundry:1.0", []string{"-v", "/Users/mac/Code/test2/src/:foundry/src"})
	//command := "cd /foundry && ls"
	//args := []string{"exec", "-i", "hello", "bash", "-c", command}
	//
	//execRes, err := execCommand("", "docker", args...)
	//fmt.Println(execRes)
	//fmt.Println(err)
	//fmt.Println(s.ExistsDocker("charming_hugle2"))
	//fmt.Println(DelDocker("hello"))
	fmt.Println(path.Join("/har", "contr.txt"))
}
