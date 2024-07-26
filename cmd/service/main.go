package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/D0K-ich/JinJi-Service/network"
	"github.com/D0K-ich/JinJi-Service/network/rest"
	"github.com/D0K-ich/JinJi-Service/store"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"

	jinji "github.com/D0K-ich/JinJi-Service"
	"github.com/D0K-ich/JinJi-Service/logs"
	"github.com/fasthttp/session/v2"
)

var (
	mainCtx, mainCancel = context.WithCancel(context.Background())
	configPath          = flag.String("config", "..\\..\\templates\\jinji.yml", "Config file path")
	config              *jinji.Config
	log                    *zap.Logger

	GitTag   string
	CommitId string
	version  string
)

func init() {
	var err error

	if config, err = jinji.NewConfig(configPath); err != nil {panic("Failed create config" + err.Error())}
	config.Print()

	if err = logs.SetConf(config.Logger); err != nil {panic("Failed create new logger" + err.Error())}
	log = logs.NewLog()

	var sig_chan = make(chan os.Signal)
	signal.Notify(sig_chan, os.Interrupt, syscall.SIGTERM)
	go func() {
		var sig = <-sig_chan
		log.Info("(main) >> Got os signal", zap.Any("signal", sig))
		mainCancel()
	}()

	version = fmt.Sprintf("%s:%s", GitTag, CommitId)
}

func main() {
	var err error

	log.Info("(main) >> Starting app...")

	log.Info("(main) >> Creating store...")
	if store.Default, err = store.NewStore(config.Store); err != nil {
		log.Fatal("Error while create store", zap.Any("error", err))
		return
	}

	log.Info("Creating new user session")
	var user_session *session.Session
	if user_session, err = rest.NewSession(config.Rest, rest.CookieNameUser, rest.TableUserSessions); err != nil {
		log.Fatal("(main) >> failed to create user session", zap.Any("err", err))
	}

	log.Info("(main) >> Creating router...")
	if network.DefaultServer, err = network.NewServer(config.Server, version, user_session); err != nil {
		log.Fatal("Err while create server", zap.Any("err", err))
	}

	log.Info("(main) >> Creating gpt...")

	log.Info("(main) >> Ready to cook")

	select {
	case <-mainCtx.Done():
		log.Info("(main) >> Shutting down...")
		return
	}
}
