package main

import (
	"time"
)

// func TestServer() {
// 	server1 := gnet.MakeServer("[test demo]")
// 	server1.Serve()
// }
// func TestClient() {
// 	fmt.Println("client start connect---")
// 	time.Sleep(1 * time.Second)

// 	connect, err := net.Dial("tcp", "127.0.0.1:12323")
// 	if err != nil {

// 	}
// 	buf := make([]byte, 512)
// 	for {
// 		_, err2 := connect.Write([]byte("hello test client"))
// 		if err2 != nil {

// 		}
// 		cnt, err3 := connect.Read(buf)
// 		if err3 != nil {

// 		}
// 		fmt.Printf("recive messege: %s length:%d\n", buf, cnt)
// 	}

// }

func main() {
	go TestServer()

	time.Sleep(3 * time.Second)

	go TestClient()

	time.Sleep(100 * time.Second)

}
