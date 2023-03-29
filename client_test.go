package client

import (
	"reflect"
	"testing"
)

func TestChatGPTRedisClient_Call(t *testing.T) {
	type args struct {
		request string
	}
	tests := []struct {
		name string
		c    *ChatGPTRedisClient
		args args
		want string
	}{
		{
			name: "case1",
			c: NewClient(
				"",
				"",
				0,
			),
			args: args{"你好"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Call(tt.args.request); got != tt.want {
				t.Errorf("ChatGPTRedisClient.Call() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		redisHost     string
		redisPassword string
		redisDB       int
	}
	tests := []struct {
		name string
		args args
		want *ChatGPTRedisClient
	}{
		{
			name: "case1",
			args: args{
				"",
				"",
				0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.redisHost, tt.args.redisPassword, tt.args.redisDB); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
