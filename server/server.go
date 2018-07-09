package server

import (
	"net"
	"log"
	"fmt"
	"io"
	"bytes"
)

func NewServer(addr string, handler func(ctx *TcportContext)) (*tcportServer, error) {
	return &tcportServer{
		handler: handler,
		addr: addr,
	}, nil
}

type tcportServer struct {
	handler func (ctx *TcportContext)
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
			log.Printf("error accepting connection %s", err)
			continue
		}
		ctx, err := s.sessionStart(conn)
		if err != nil {
			log.Printf("error accepting session %s", err)
			conn.Close()
			continue
		}
		log.Printf("Session started")
		s.handler(ctx)
	}
}

func (s *tcportServer) sessionStart(conn net.Conn) (*TcportContext, error) {
	buf := &bytes.Buffer{}
	_, err := io.CopyN(buf, conn, 4)
	if err != nil {
		return nil, fmt.Errorf("error of reading session headers")
	}

	if buf.Len() < 4 {
		return nil, fmt.Errorf("invalid session syn")
	}

	if buf.Bytes()[0] != '0' {
		return nil, fmt.Errorf("unexpected session version")
	}

	if buf.Bytes()[1] != '0' {
		return nil, fmt.Errorf("unexpected content type")
	}

	return &TcportContext{
		Conn: conn,
		ContentType: "application/json",
	}, nil
}
