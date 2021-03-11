package proc

import (
	"sync"

	"github.com/iegad/kraken/nw/server"
)

var poolUserData = &sync.Pool{
	New: func() interface{} {
		return &userData{}
	},
}

type userData struct {
	conn server.IConn
}

func (this_ *userData) set(conn server.IConn) {
	this_.conn = conn
}

func (this_ *userData) reset() {
	this_.conn.Reset()
}
