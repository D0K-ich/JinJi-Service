package gpt

import (
	"context"
	"fmt"
	//"github.com/D0K-ich/JinJi-Service/logs"
	openai "github.com/sashabaranov/go-openai"
	"net/http"
	"net/url"
	"testing"
)

//func TestGpt(t *testing.T) {
//	var err error
//	if err = logs.SetConf(&logs.Config{
//		Level:  "debug",
//		Output: "console",
//	}); err != nil {
//		panic("Failed create new logger" + err.Error())
//	}
//	NewDefaultGPT(context.Background(), nil)
//}

func TestReqGpt(t *testing.T) {
	var err error

	var config = openai.DefaultConfig("sk-proj-1hPIEis3bXZ943oWsR8bT3BlbkFJ01XOOY0kyv9tqi6xfQM5")

	var proxyURL, _ = url.Parse("http://51.161.56.52:80")
	var proxy = http.ProxyURL(proxyURL)
	var transport = &http.Transport{Proxy: proxy}

	config.HTTPClient = &http.Client{Transport: transport}
	client := openai.NewClientWithConfig(config)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)

}

//sk-proj-1hPIEis3bXZ943oWsR8bT3BlbkFJ01XOOY0kyv9tqi6xfQM5
