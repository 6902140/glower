package router

import (
	"github.com/6902140/glower/ziface"
	"github.com/6902140/glower/zlog"
	"github.com/6902140/glower/znet"
)

type HelloRouter struct {
	znet.BaseRouter
}

func (hr *HelloRouter) Handle(request ziface.IRequest) {
	zlog.Ins().InfoF(string(request.GetData()))
}
