package main

import (
	"github.com/iegad/cerberus/cfg"
	"github.com/iegad/cerberus/proc"
	"github.com/iegad/cerberus/proc/handlers"
	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/nw"
	"github.com/iegad/kraken/nw/server"
)

func main() {
	err := cfg.Init("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	cerberus := proc.NewCerberus()

	cerberus.Regist(&handlers.UserSignIn{})
	cerberus.Regist(&handlers.UserSignOut{})
	cerberus.Regist(&handlers.Delivery{})

	opt := &server.Option{
		Host:    cfg.Instance.Server.Host,
		MaxConn: cfg.Instance.Server.MaxConn,
		Timeout: cfg.Instance.Server.Timeout,
	}

	server, err := nw.NewServer(cfg.Instance.Server.Protocol, cerberus, opt)
	if err != nil {
		log.Fatal(err)
	}

	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
