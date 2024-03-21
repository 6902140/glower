package c_router

import (
	"github.com/6902140/glower/ziface"
	"github.com/6902140/glower/zlog"
	"github.com/6902140/glower/znet"
)

// Ping test custom routing.
type PingRouter struct {
	znet.BaseRouter
}

// Ping Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	zlog.Debug("Call PingRouter Handle")
	zlog.Debug("recv from server : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	if err := request.GetConnection().SendBuffMsg(1, []byte("Hello[from client]")); err != nil {
		zlog.Error(err)
	}
}
