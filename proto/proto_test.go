package proto

import (
	"github.com/D0K-ich/JinJi-Service/proto/neo"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	var err error
	if err = neo.StartgRPC(); err != nil {t.Error(err)}

	time.Sleep(60 * time.Hour)
}
