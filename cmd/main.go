package main

import (
	"fmt"
	"net"
)

func main() {
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
		fmt.Println(string(body))

		ParseFirstLine(string(body))
	}
}
