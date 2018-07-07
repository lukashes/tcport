package main

import (
	"net"
	"log"
	"bufio"

	"github.com/lukashes/tcport/server"
)

func main() {
	addr := "localhost:8080"
	s, err := server.NewServer(addr, handleRequest)
	if err != nil {
		log.Fatalf("New server: %s", err.Error())
	}
	log.Printf("Start listening at: %s", addr)
	log.Fatalf("Listening: %s", s.ListenAndServe())
}

func handleRequest(conn net.Conn) {
	r := bufio.NewReader(conn)
	scan := bufio.NewScanner(r)

	for scan.Scan() {
		if err := scan.Err(); err != nil {
			log.Printf("Read request: %s", err)
			conn.Write([]byte("Request error"))
			break
		}
		conn.Write(append([]byte("Received request: "), scan.Bytes()...))
		conn.Write([]byte("\n"))
	}

	conn.Close()
}