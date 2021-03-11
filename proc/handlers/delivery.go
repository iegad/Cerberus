package handlers

import (
	"github.com/iegad/cerberus/proc"
	"github.com/iegad/kraken/pb"
)

type Delivery struct {
}

func (this_ *Delivery) PID() int32 {
	return 0
}

func (this_ *Delivery) Do(ctx *proc.Context, pack *pb.Package) error {
	// TODO: 消息转发
	return nil
}
