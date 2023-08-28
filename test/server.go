package main

import "github.com/Okkotsu/glower/gnet"

func TestServer() {
	server1 := gnet.MakeServer("[test demo]")
	server1.Serve()
}
