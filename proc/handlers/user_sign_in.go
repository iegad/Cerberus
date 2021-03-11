package handlers

import (
	"github.com/iegad/cerberus/proc"
	"github.com/iegad/kraken/pb"
)

type UserSignIn struct {
}

func (this_ *UserSignIn) PID() int32 {
	return 0
}

func (this_ *UserSignIn) Do(ctx *proc.Context, pack *pb.Package) error {
	// TODO: 用户登录
	return nil
}
