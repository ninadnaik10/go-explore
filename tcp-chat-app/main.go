package main

import (
	"bufio"
	"fmt"
	"net"
)

var Client struct {
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Accepted from: ", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		txt := scanner.Text()
		fmt.Printf("recv: %s\n", txt)
		_, _ = conn.Write([]byte("echo: " + txt + "\n"))
		if txt == "quit" {
			return
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	defer ln.Close()

	fmt.Println("listening on :8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Accept error: ", err)
			continue
		}

		go handleConnection(conn)

	}
}
