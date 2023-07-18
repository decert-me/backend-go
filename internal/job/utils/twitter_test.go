package utils

import (
	"backend-go/internal/job/config"
	"backend-go/internal/job/initialize"
	"testing"
)

var c *config.Config

func TestMain(m *testing.M) {
	c = initialize.Viper("../cmd/config.yaml")
	m.Run()
}

func TestGetSpyderTweetById(t *testing.T) {
	GetSpyderTweetById(c, "1664476729812094977")

}
