package proc

import "github.com/iegad/kraken/pb"

type IHandler interface {
	PID() int32
	Do(ctx *Context, pack *pb.Package) error
}
