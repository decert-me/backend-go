package service

import (
	"backend-go/internal/judge/initialize"
	"backend-go/pkg/log"
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
	//CreateDocker("hello", "foundry:1.0", []string{"-v", "/Users/mac/Code/test2/src/:foundry/src"})
	//fmt.Println(s.ExistsDocker("charming_hugle2"))
	//fmt.Println(DelDocker("hello"))
}
