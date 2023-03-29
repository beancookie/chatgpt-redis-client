package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bwmarrin/snowflake"
	redis "github.com/redis/go-redis/v9"
)

const CHATGPT_CHANNEL = "chatGPT"

type Message struct {
	ChannelId string
	Request   string
}

type ChatGPTRedisClient struct {
	rdb   *redis.Client
	snode *snowflake.Node
}

func (c ChatGPTRedisClient) Call(request string) string {
	channelId := c.snode.Generate()

	// 定义一个Message结构体实例
	msg := Message{
		ChannelId: fmt.Sprintf("%d", channelId),
		Request:   request,
	}

	// 将Message结构体序列化为JSON格式的字符串
	requestJson, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Json encode error:", err)
		panic(err)
	}

	err = c.rdb.Publish(context.Background(), CHATGPT_CHANNEL, requestJson).Err()
	if err != nil {
		fmt.Println("Publish error:", err)
		panic(err)
	}

	pubsub := c.rdb.Subscribe(context.Background(), msg.ChannelId)
	defer pubsub.Close()

	resp := <-pubsub.Channel()
	return resp.Payload
}

func NewClient(redisHost string, redisPassword string, redisDB int) *ChatGPTRedisClient {
	snode, _ := snowflake.NewNode(1)

	c := ChatGPTRedisClient{
		rdb: redis.NewClient(&redis.Options{
			Addr:     redisHost,
			Password: redisPassword,
			DB:       redisDB,
		}),
		snode: snode,
	}

	return &c
}
