package proc

import (
	"sync"
	"sync/atomic"

	"github.com/iegad/cerberus/cfg"
	"github.com/iegad/kraken/cs"
	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/nw/server"
	"github.com/iegad/kraken/pb"
	"google.golang.org/protobuf/proto"
)

var poolPack = sync.Pool{
	New: func() interface{} {
		return &pb.Package{}
	},
}

var poolContent = sync.Pool{
	New: func() interface{} {
		return &Context{}
	},
}

type Context struct {
	Server   server.IServer
	UserData *UserData
}

type Cerberus struct {
	userMap   sync.Map
	handlers  map[int32]IHandler
	connCount int32
	userCount int32
}

func NewCerberus() *Cerberus {
	return &Cerberus{
		handlers: make(map[int32]IHandler),
	}
}

func (this_ *Cerberus) OnInit(svr server.IServer) {
	err := cs.Regist(&cs.Option{
		Host:    cfg.Instance.Consul.Host,
		Service: cfg.Instance.Consul.Service,
		ID:      cfg.Instance.Consul.ID,
		Addr:    cfg.Instance.Server.Host,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (this_ *Cerberus) Process(svr server.IServer, conn server.IConn, data []byte) bool {
	var (
		ctx        = poolContent.Get().(*Context)
		pack       = poolPack.Get().(*pb.Package)
		found, res = false, false
		h          IHandler
	)

	ctx.UserData = this_.getUser(conn)

	err := proto.Unmarshal(data, pack)
	if err != nil {
		log.Error(err)
		goto PROC_EXIT
	}

	h, found = this_.handlers[pack.PID]
	if !found {
		log.Error("PID[%d] is invalid", pack.PID)
		goto PROC_EXIT
	}

	err = h.Do(ctx, pack)
	if err != nil {
		log.Error(err)
		goto PROC_EXIT
	}

	poolContent.Put(ctx)
	poolPack.Put(pack)

	res = true

PROC_EXIT:
	return res
}

func (this_ *Cerberus) OnConnected(svr server.IServer, conn server.IConn) {
	atomic.AddInt32(&this_.connCount, 1)
	this_.addUser(conn)
}

func (this_ *Cerberus) OnDisconnect(svr server.IServer, conn server.IConn) {
	ud := this_.getUser(conn)
	if ud.UserID > 0 {
		atomic.AddInt32(&this_.userCount, -1)
	}

	this_.rmvUser(conn)
	atomic.AddInt32(&this_.connCount, -1)
}

func (this_ *Cerberus) Regist(h IHandler) {
	if _, found := this_.handlers[h.PID()]; found {
		log.Fatal("%d is already exists", h.PID())
	}

	this_.handlers[h.PID()] = h
}

func (this_ *Cerberus) addUser(conn server.IConn) {
	ud := this_.getUser(conn)
	if ud != nil {
		ud.reset()
	} else {
		ud = poolUserData.Get().(*UserData)
	}

	ud.set(conn)
}

func (this_ *Cerberus) getUser(conn server.IConn) *UserData {
	v, found := this_.userMap.Load(conn.RouteKey())
	if found {
		return v.(*UserData)
	}

	return nil
}

func (this_ *Cerberus) rmvUser(conn server.IConn) {
	v, found := this_.userMap.LoadAndDelete(conn.RouteKey())
	if found {
		ud := v.(*UserData)
		if ud != nil {
			poolUserData.Put(ud)
		}
	}
}
