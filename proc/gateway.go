package proc

import (
	"github.com/iegad/cerberus/cfg"
	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/nw"
	"github.com/iegad/kraken/nw/server"
)

type Gateway struct {
	fsvr server.IServer
}

func NewGateway(config *cfg.Config) *Gateway {
	var (
		this_ = &Gateway{}
		err   error
	)

	this_.fsvr, err = nw.NewServer(
		config.Server.Protocol,
		newFEngine(
			config.Consul.Host,
			config.Server.Host,
			config.Consul.Service,
			config.Consul.ID,
		),
		&server.Option{
			Host:    config.Server.Host,
			MaxConn: config.Server.MaxConn,
			Timeout: config.Server.Timeout,
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	return this_
}

func (this_ *Gateway) Run() {
	this_.fsvr.Run()
}
