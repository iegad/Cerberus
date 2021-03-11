package proc

import (
	"encoding/json"
	"sync"

	"github.com/iegad/kraken/nw/server"
)

var poolUserData = &sync.Pool{
	New: func() interface{} {
		return &UserData{}
	},
}

type UserData struct {
	conn   server.IConn
	UserID int64 `json:"userID"`
}

func (this_ *UserData) String() string {
	data, _ := json.Marshal(this_)
	return string(data)
}

func (this_ *UserData) set(conn server.IConn) {
	this_.conn = conn
}

func (this_ *UserData) reset() {
	this_.conn.Reset()
}
