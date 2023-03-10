package utils

import (
	"backend-go/internal/app/config"
	"backend-go/internal/app/initialize"
	"testing"
)

var c *config.Config

func TestMain(m *testing.M) {
	c = initialize.Viper("../cmd/config.yaml")
	m.Run()
}

func Test_getTweetIdFromURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name        string
		args        args
		wantTweetId string
	}{
		{name: "#0", args: args{url: "https://twitter.com/NUpchain/status/1626415377281593345?s=20"}, wantTweetId: "1626415377281593345"},
		{name: "#1", args: args{url: "https://twitter.com/taylorswift13/status/1587420273325838336"}, wantTweetId: "1587420273325838336"},
		{name: "#2", args: args{url: "https://twitter.com/nytchinese/status/1626748330611974147?https://twitter.com/nytchinese/status/1626748330611974147"}, wantTweetId: "1626748330611974147"},
		{name: "#3", args: args{url: ""}, wantTweetId: ""},
		{name: "#4", args: args{url: "https://twitter.com/notifications"}, wantTweetId: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTweetId := GetTweetIdFromURL(tt.args.url); gotTweetId != tt.wantTweetId {
				t.Errorf("getTweetIdFromURL() = %v, want %v", gotTweetId, tt.wantTweetId)
			}
		})
	}
}
