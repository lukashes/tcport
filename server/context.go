package server

import "net"

type TcportContext struct {
	Conn        net.Conn
	ContentType string
}
