package main

import (
	"fmt"
	"github.com/6902140/glower/ziface"
	"github.com/6902140/glower/znet"
	"time"
)

func DefaultTest1(request ziface.IRequest) {
	fmt.Println("test1")
}
func DefaultTest2(request ziface.IRequest) {
	time.Sleep(1)
	arr := make([]int, 1)
	fmt.Println(arr[1])
}

func main() {
	s := znet.NewDefaultRouterSlicesServer()
	s.AddRouterSlices(1, DefaultTest1, DefaultTest2)
	s.Serve()
}
