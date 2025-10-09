package main

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"
)

var (
	peers = make(map[string]bool)
	mu    sync.Mutex
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	var addr string
	fmt.Fscanln(conn, &addr)

	mu.Lock()
	peers[addr] = true
	var list []string
	for p := range peers {
		if p != addr {
			list = append(list, p)
		}
	}
	mu.Unlock()

	data, _ := json.Marshal(list)
	conn.Write(data)
}

func main() {
	ln, _ := net.Listen("tcp", ":7000")
	fmt.Println("Discovery server on :7000")

	for {
		conn, _ := ln.Accept()
		go handleConn(conn)
	}
}
