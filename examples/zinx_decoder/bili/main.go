package main

import (
	"github.com/6902140/glower/examples/zinx_decoder/bili/router"
	"github.com/6902140/glower/zdecoder"
	"github.com/6902140/glower/ziface"
	"github.com/6902140/glower/znet"
)

func DoConnectionBegin(conn ziface.IConnection) {
}

func DoConnectionLost(conn ziface.IConnection) {
}

func main() {
	server := znet.NewServer(func(s *znet.Server) {
		s.Port = 9090
		/*
			s.LengthField = ziface.LengthField{
				MaxFrameLength:      math.MaxUint8 + 4,
				LengthFieldOffset:   2,
				LengthFieldLength:   1,
				LengthAdjustment:    2,
				InitialBytesToStrip: 0,
			}
		*/
	})
	server.SetOnConnStart(DoConnectionBegin)
	server.SetOnConnStop(DoConnectionLost)
	server.AddInterceptor(zdecoder.NewHTLVCRCDecoder())
	server.AddRouter(0x10, &router.Data0x10Router{})
	server.AddRouter(0x13, &router.Data0x13Router{})
	server.AddRouter(0x14, &router.Data0x14Router{})
	server.AddRouter(0x15, &router.Data0x15Router{})
	server.AddRouter(0x16, &router.Data0x16Router{})
	server.Serve()
}
