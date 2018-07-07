package server

import (
	"net"
	"log"
)

func NewServer(addr string, handler func(conn net.Conn)) (*tcportServer, error) {
	return &tcportServer{
		handler: handler,
		addr: addr,
	}, nil
}

type tcportServer struct {
	handler func (conn net.Conn)
	addr    string
}

func (s *tcportServer) ListenAndServe() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("error accepting connection %v", err)
			continue
		}
		log.Printf("accepted connection from %v", conn.RemoteAddr())
		s.handler(conn)
	}
}
