package server

import (
	"fmt"
	"net"
	"sla/internal/cmd"
)

func Start() {
	ln, _ := net.Listen("tcp", "0.0.0.0:3030")
	defer ln.Close()

	conn, _ := ln.Accept()
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Cliente desconectou")
			return
		}

		output, err := cmd.HandleCommands(buf[:n])
		conn.Write([]byte(output))

	}
}
