package gnet

import (
	"fmt"
	"net"

	"github.com/Okkotsu/glower/serverFace"
)

type Server struct {
	Version string
	IP      string
	Port    int
	Name    string
}

func (s Server) Start() {
	tcpAddr, err := net.ResolveTCPAddr(s.Version, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("failed to resolve the TCP adders")
		return
	}
	listener, err := net.ListenTCP(s.Version, tcpAddr)
	if err != nil {
		fmt.Println("failed to listen at TCPaddr")
		return
	}

	for {
		fmt.Println("here to wait for accept")
		conn, err := listener.Accept()

		fmt.Printf("\033[1;34;40m%s\033[0m\n", "accept success")

		if err != nil {
			continue
		}
		go HandlerServer(conn)

	}

}
func (s Server) Serve() {
	s.Start()

	select {}
}

func (s Server) Close() {

}

func HandlerServer(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	count, err := conn.Read(buf)
	if err != nil {
		return
	} else {
		_, err := conn.Write(buf[:count])
		fmt.Printf("\033[1;36;40mserver\033[0m start return back :%s ;\n", buf)
		if err != nil {
			return
		}
	}

}

func MakeServer(name string) serverFace.Serverf {
	res := &Server{
		Version: "tcp4",
		IP:      "0.0.0.0",
		Name:    name,
		Port:    12323,
	}
	return res
}
