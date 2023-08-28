package main

import (
	"fmt"
	"net"
	"time"
)

func TestClient() {
	fmt.Println("client start connect---")
	time.Sleep(1 * time.Second)

	buf := make([]byte, 512)
	for {
		connect, _ := net.Dial("tcp", "127.0.0.1:12323")
		time.Sleep(1 * time.Second)
		_, err2 := connect.Write([]byte("\033[1;32;40mclient\033[0m hello test " + time.Now().Format(time.DateTime)))
		time.Sleep(1 * time.Second)
		if err2 != nil {
			fmt.Println("connect write err")
		}
		cnt, err3 := connect.Read(buf)
		if err3 != nil {
			fmt.Println("connect read err")
		}
		fmt.Printf("\033[1;34;40mclient\033[0m recive messege: %s length:%d\n", buf, cnt)
	}

}
