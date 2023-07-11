package dao

import (
	"backend-go/internal/auth/initialize"
	"backend-go/pkg/log"
	"fmt"
	"testing"
)

func TestDao_TwitterBinding(t *testing.T) {
	c := initialize.Viper("../../../bin/auth/config.yaml")
	log.Init(c.Log)
	d := New(c)
	//fmt.Println(d.TwitterIsBinding(123421))
	//d.TwitterBinding("0xBC5Ea980BdD0436a2798Bccf8fECc61bCb0010f2", 1232, "liangejis")
	fmt.Println(d.TwitterQueryIdByAddress("0xBC5Ea980BdD0436a2798Bccf8fECc61bCb0010f2"))
}
