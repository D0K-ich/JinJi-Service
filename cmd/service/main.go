package main

import (
	"os"
	"fmt"
	"flag"
	"context"
	"syscall"
	"os/signal"

	"github.com/rs/zerolog/log"
	"github.com/fasthttp/session/v2"

	"github.com/D0K-ich/JinJi-Service/store"
	"github.com/D0K-ich/JinJi-Service/logs"
	jinji "github.com/D0K-ich/JinJi-Service"
	"github.com/D0K-ich/JinJi-Service/network"
	"github.com/D0K-ich/JinJi-Service/network/rest"
)

var (
	mainCtx, mainCancel = context.WithCancel(context.Background())
	configPath          = flag.String("config", "..\\..\\templates\\jinji.yml", "Config file path")
	config              *jinji.Config

	GitTag   string
	CommitId string
	version  string
)

func init() {
	var err error

	if config, err = jinji.NewConfig(configPath); err != nil {panic("Failed create config" + err.Error())}
	config.Print()

	if err = logs.SetConf(config.Logger); err != nil {panic("Failed create new logger" + err.Error())}

	var sig_chan = make(chan os.Signal)
	signal.Notify(sig_chan, os.Interrupt, syscall.SIGTERM)
	go func() {
		var sig = <-sig_chan
		log.Info().Msgf("(main) >> Got os signal %s %s", "signal", sig)
		mainCancel()
	}()

	version = fmt.Sprintf("%s:%s", GitTag, CommitId)
}

func main() {
	var err error

	log.Info().Msg("(main) >> Starting app...")

	log.Info().Msg("(main) >> Creating store...")
	if store.Default, err = store.NewStore(config.Store); err != nil {
		log.Fatal().Msgf("Error while create store", "error", err)
		return
	}

	log.Info().Msg("(main) >> Creating new user session...")
	var user_session *session.Session
	if user_session, err = rest.NewSession(config.Rest, rest.CookieNameUser, rest.TableUserSessions); err != nil {
		log.Fatal().Msgf("(main) >> failed to create user session", "err", err)
	}

	log.Info().Msg("(main) >> Creating router...")
	if network.DefaultServer, err = network.NewServer(config.Server, version, user_session); err != nil {
		log.Fatal().Msgf("(main) >> Err while create server", "err", err)
	}

	log.Info().Msg("(main) >> Creating gpt...")

	log.Info().Msg("(main) >> Ready to cook")

	select {
	case <-mainCtx.Done():
		log.Warn().Msg("(main) >> Shutting down...")
		return
	}
}
