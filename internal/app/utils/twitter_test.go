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

func Test_checkIfMatchClaimTweet(t *testing.T) {
	c.Twitter.ClaimContent = "我通过了@DecertMe的挑战并获得了一个链上的能力认证徽章。\nhttps://decert.me/quests/\n#Decert.me"
	type args struct {
		tokenId int64
		tweet   string
	}
	tests := []struct {
		name     string
		args     args
		wantPass bool
	}{
		{name: "#1 should pass", args: args{tokenId: 10006, tweet: "我通过了@DecertMe的挑战并获得了一个链上的能力认证徽章。\nhttps://t.co/gMf1CuE8pS\n#Decert.me"}, wantPass: true},
		{name: "#2 should fail tokenID not equal", args: args{tokenId: 10001, tweet: "我通过了@DecertMe的挑战并获得了一个链上的能力认证徽章。\nhttps://t.co/gMf1CuE8pS\n#Decert.me"}, wantPass: false},
		{name: "#3 should fail text not complete", args: args{tokenId: 10001, tweet: "我通过了@DecertMe的挑战并获得了一个链上的能力认证徽章。\nhttps://t.co/gMf1CuE8pS\n"}, wantPass: false},
		{name: "#4 should fail empty text", args: args{tokenId: 10001, tweet: ""}, wantPass: false},
		{name: "#5 should fail wrong link", args: args{tokenId: 10001, tweet: "我通过了@DecertMe的挑战并获得了一个链上的能力认证徽章。\nhttps://learnblockchain.cn/article/5435\n#Decert.me"}, wantPass: false},
		{name: "#6 should pass has bank", args: args{tokenId: 10006, tweet: "我通过了@DecertMe的挑战并获得了一个链上的能力认证徽章。\n https://t.co/gMf1CuE8pS\n #Decert.me"}, wantPass: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPass := CheckIfMatchClaimTweet(c, tt.args.tokenId, tt.args.tweet); gotPass != tt.wantPass {
				t.Errorf("checkIfMatchClaimTweet() = %v, want %v", gotPass, tt.wantPass)
			}
		})
	}
}

func TestGetTweetById(t *testing.T) {
	c.Twitter.ClaimContent = "我通过了@DecertMe的挑战并获得了一个链上的能力认证徽章。\nhttps://decert.me/quests/\n#Decert.me"
	type args struct {
		c       *config.Config
		tweetId string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "#0 should return right data", args: args{c: c, tweetId: "1633028821715927041"}, want: "我在 @DecertMe 上完成了一个挑战并获得了链上能力认证的徽章。\n https://t.co/YkPclYsZYw\n #DecertMe"},
		{name: "#1 should return empty data when tweeId no correct", args: args{c: c, tweetId: "123124214"}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTweetById(tt.args.c, tt.args.tweetId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTweetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetTweetById() got = %v, want %v", got, tt.want)
			}
		})
	}
}
