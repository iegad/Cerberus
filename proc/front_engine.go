package proc

import (
	"sync"

	"github.com/iegad/kraken/cs"
	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/nw/server"
)

type frontEngine struct {
	userMap sync.Map
	host    string
	addr    string
	service string
	id      string
}

func newFEngine(host, addr, service, id string) *frontEngine {
	return &frontEngine{
		host:    host,
		addr:    addr,
		service: service,
		id:      id,
	}
}

func (this_ *frontEngine) OnInit(svr server.IServer) {
	err := cs.Regist(&cs.Option{
		Host:    this_.host,
		Service: this_.service,
		ID:      this_.id,
		Addr:    this_.addr,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (this_ *frontEngine) Process(svr server.IServer, conn server.IConn, data []byte) bool {
	ud := this_.getUser(conn)

	log.Info(ud)

	return true
}

func (this_ *frontEngine) OnConnected(svr server.IServer, conn server.IConn) {
	this_.addUser(conn)
}

func (this_ *frontEngine) OnDisconnect(svr server.IServer, conn server.IConn) {
	this_.rmvUser(conn)
}

func (this_ *frontEngine) addUser(conn server.IConn) {
	ud := this_.getUser(conn)
	if ud != nil {
		ud.reset()
	} else {
		ud = poolUserData.Get().(*userData)
	}

	ud.set(conn)
}

func (this_ *frontEngine) getUser(conn server.IConn) *userData {
	v, found := this_.userMap.Load(conn.RouteKey())
	if found {
		return v.(*userData)
	}

	return nil
}

func (this_ *frontEngine) rmvUser(conn server.IConn) {
	v, found := this_.userMap.LoadAndDelete(conn.RouteKey())
	if found {
		ud := v.(*userData)
		if ud != nil {
			poolUserData.Put(ud)
		}
	}
}
