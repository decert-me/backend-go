package utils

import (
	"backend-go/internal/app/config"
	"testing"
)

var c config.Config

func TestMain(m *testing.M) {
	c = config.Config{Twitter: &config.Twitter{ClaimContent: "我通过了@DecertMe的挑战并获得了一个链上的能力认证徽章。\nhttps://decert.me/quests/\n#Decert.me"}}
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

func Test_checkIfMatchClaimTweet(t *testing.T) {
	type args struct {
		tokenId int64
		tweet   string
	}
	tests := []struct {
		name     string
		args     args
		wantPass bool
	}{
		{name: "#1", args: args{tokenId: 10006, tweet: "我通过了@DecertMe的挑战并获得了一个链上的能力认证徽章。\nhttps://t.co/gMf1CuE8pS\n#Decert.me"}, wantPass: true},
		{name: "#2", args: args{tokenId: 10001, tweet: "我通过了@DecertMe的挑战并获得了一个链上的能力认证徽章。\nhttps://t.co/gMf1CuE8pS\n#Decert.me"}, wantPass: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPass := CheckIfMatchClaimTweet(&c, tt.args.tokenId, tt.args.tweet); gotPass != tt.wantPass {
				t.Errorf("checkIfMatchClaimTweet() = %v, want %v", gotPass, tt.wantPass)
			}
		})
	}
}
