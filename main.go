package main

import (
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

func handleRequest(ctx *server.TcportContext) {
	r := bufio.NewReader(ctx.Conn)
	scan := bufio.NewScanner(r)

	for scan.Scan() {
		if err := scan.Err(); err != nil {
			log.Printf("Read request: %s", err)
			ctx.Conn.Write([]byte("Request error"))
			break
		}
		ctx.Conn.Write(append([]byte("Received request: "), scan.Bytes()...))
		ctx.Conn.Write([]byte("\n"))
	}

	ctx.Conn.Close()
}