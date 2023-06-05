package utils

import (
	"fmt"
	"testing"
)

func TestTT(t *testing.T) {
	fmt.Println(AnswerDecode("Pz8HN1LAD7Q2C6gvQewnM3pxFvLeO8z8eKNnWxGQ", "C0oUE34dfRwZ"))
	fmt.Println(AnswerEncode("Pz8HN1LAD7Q2C6gvQewnM3pxFvLeO8z8eKNnWxGQ", "[0,[0,1],\"true\"]"))
}

/*
func TestAnswerCheck(t *testing.T) {
	type args struct {
		key        string
		answerUser string
		uri        string
		score      int64
	}
	tests := []struct {
		name     string
		args     args
		wantPass bool
		wantErr  bool
	}{
		{name: "#0 should pass score 40/40", args: args{key: "eb5a5bb2-ebbd-45cc-9d37-77a9377f2aca", answerUser: "[[0,33],\"答案\",\"答案2\",0]", uri: "ipfs://QmXuPifE21Jo3uReKoDRF2ctCLferCbgBL7F9GzmrhNhF4", score: 10000}, wantPass: true},
		{name: "#1 should pass score 30/40", args: args{key: "eb5a5bb2-ebbd-45cc-9d37-77a9377f2aca", answerUser: "[[0,33],\"答案\",\"答案2\",1]", uri: "ipfs://QmXuPifE21Jo3uReKoDRF2ctCLferCbgBL7F9GzmrhNhF4", score: 7500}, wantPass: true},
		{name: "#2 should fail pass score 20/40", args: args{key: "eb5a5bb2-ebbd-45cc-9d37-77a9377f2aca", answerUser: "[[0,33],\"答案\",\"答案1\",1]", uri: "ipfs://QmXuPifE21Jo3uReKoDRF2ctCLferCbgBL7F9GzmrhNhF4", score: 5000}, wantPass: false},
		{name: "#3 should fail pass score 10/40", args: args{key: "eb5a5bb2-ebbd-45cc-9d37-77a9377f2aca", answerUser: "[[1,33],\"答案\",\"答案1\",1]", uri: "ipfs://QmXuPifE21Jo3uReKoDRF2ctCLferCbgBL7F9GzmrhNhF4", score: 5000}, wantPass: false},
		{name: "#4 should fail pass score 0/40", args: args{key: "eb5a5bb2-ebbd-45cc-9d37-77a9377f2aca", answerUser: "[[1,33],\"答案3\",\"答案1\",1]", uri: "ipfs://QmXuPifE21Jo3uReKoDRF2ctCLferCbgBL7F9GzmrhNhF4", score: 10000}, wantPass: false},
		{name: "#5 should fail not enough quantity", args: args{key: "eb5a5bb2-ebbd-45cc-9d37-77a9377f2aca", answerUser: "[[1,33],\"答案3\",\"答案1\"]", uri: "ipfs://QmXuPifE21Jo3uReKoDRF2ctCLferCbgBL7F9GzmrhNhF4", score: 10000}, wantErr: true},
		{name: "#5 should fail uri invalid", args: args{key: "eb5a5bb2-ebbd-45cc-9d37-77a9377f2aca", answerUser: "[[0,33],\"答案\",\"答案2\",0]", uri: "ipfs://QmXuPi", score: 10000}, wantErr: true},
		{name: "#6 should fail score 30/40 then higher", args: args{key: "eb5a5bb2-ebbd-45cc-9d37-77a9377f2aca", answerUser: "[[0,33],\"答案\",\"答案2\",1]", uri: "ipfs://QmXuPifE21Jo3uReKoDRF2ctCLferCbgBL7F9GzmrhNhF4", score: 7501}, wantPass: false},
		{name: "#7 should fail score 30/40 then slower", args: args{key: "eb5a5bb2-ebbd-45cc-9d37-77a9377f2aca", answerUser: "[[0,33],\"答案\",\"答案2\",1]", uri: "ipfs://QmXuPifE21Jo3uReKoDRF2ctCLferCbgBL7F9GzmrhNhF4", score: 7499}, wantPass: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPass, err := AnswerCheck(tt.args.key, tt.args.answerUser, tt.args.uri, tt.args.score)
			if (err != nil) != tt.wantErr {
				t.Errorf("AnswerCheck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPass != tt.wantPass {
				t.Errorf("AnswerCheck() gotPass = %v, want %v", gotPass, tt.wantPass)
			}
		})
	}
}

func Test_answerEncode(t *testing.T) {
	type args struct {
		key  string
		data string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "#0", args: args{key: "key", data: "hello"}},
		{name: "#1", args: args{key: "", data: ""}},
		{name: "#2", args: args{key: "", data: "hello"}, wantErr: true},
		{name: "#3", args: args{key: "12", data: "[[0,1],\"test\"]"}},
		{name: "#4", args: args{key: "95311swn37b_4n191pv_k6j40000gn", data: "[[0,1],\"test\"]"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := answerEncode(tt.args.key, tt.args.data)
			if tt.args.data != answerDecode(tt.args.key, got) {
				if tt.wantErr == false {
					t.Errorf("answerEncode() key = %v, data =  %v", tt.args.key, tt.args.data)
				}
			}
		})
	}
}
*/
