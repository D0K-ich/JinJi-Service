package gpt

import (
	"context"
	"github.com/D0K-ich/KanopyService/logs"
	"testing"
)

func TestGpt(t *testing.T) {
	var err error
	if err 	= logs.SetConf(&logs.Config{
		Level:  "debug",
		Output: nil,
	})	; err != nil {panic("Failed create new logger" + err.Error())}
	NewDefaultGPT(context.Background(), nil)
}
