package main

import (
	"fmt"
	"net"
)

func main() {

	buff := make([]byte, 1024)

	ln, err := net.Listen("tcp", "0.0.0.0:3030")
	if err != nil {
		fmt.Printf("Erro ao estabelecer conexao, %s", err.Error())
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		for {
			n, err := conn.Read(buff)
			if err != nil {
				fmt.Println("Cliente saiu")
				conn.Close()
				break
			}
			read := buff[:n]
			fmt.Println(conn.RemoteAddr(), string(read))
		}

	}

}
