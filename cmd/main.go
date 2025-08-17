package main

import (
	"fmt"
	"net"

	"github.com/Qu-Ack/tcptohttp/internal/services/parser"
)

func main() {

	parse := parser.NewParser("http/1.1")

	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("error while accepting connection : %s", err.Error())
		}

		addr := conn.RemoteAddr()
		fmt.Println(addr.String())

		body := make([]byte, 65535)
		_, err = conn.Read(body)

		if err != nil {
			fmt.Println("error while reading the connection body")
		}
		request := parse.ParseRequest(body)
		parse.PrintRequest(request)
	}
}
