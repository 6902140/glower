package main

import (
	"github.com/6902140/glower/examples/zinx_interceptor/interceptors"
	"github.com/6902140/glower/examples/zinx_interceptor/router"
	"github.com/6902140/glower/znet"
)

func main() {
	server := znet.NewServer()

	server.AddRouter(1, &router.HelloRouter{})

	// Add Custom Interceptor
	server.AddInterceptor(&interceptors.MyInterceptor{})

	server.Serve()
}
