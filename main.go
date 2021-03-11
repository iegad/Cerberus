package main

import (
	"github.com/iegad/cerberus/cfg"
	"github.com/iegad/cerberus/proc"
	"github.com/iegad/kraken/log"
)

func main() {
	config, err := cfg.New("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	gw := proc.NewGateway(config)
	gw.Run()
}
