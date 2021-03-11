package handlers

import (
	"github.com/iegad/cerberus/proc"
	"github.com/iegad/kraken/pb"
)

type UserSignOut struct {
}

func (this_ *UserSignOut) PID() int32 {
	return 0
}

func (this_ *UserSignOut) Do(ctx *proc.Context, pack *pb.Package) error {
	// TODO: 用户登出
	return nil
}
