package main

import (
	"fmt"
	"net"
)

func main() {

	conncetions := make(map[string]net.Conn)

	ln, err := net.Listen("tcp", "192.168.2.143:3030")
	if err != nil {
		fmt.Printf("Erro ao estabelecer conexao, %s", err.Error())
	}

	for {
		conn, err := ln.Accept()
		conncetions[conn.RemoteAddr().String()] = conn
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn, conncetions)

	}

}

func handleConnection(conn net.Conn, connections map[string]net.Conn) {
	defer conn.Close()

	for {
		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Println("Cliente desconectou:", conn.RemoteAddr())
			return
		}
		msg := string(buff[:n])
		fmt.Printf("[Client %s]: %s\n", conn.RemoteAddr(), msg)

		for _, connection := range connections {
			if connection == conn {
				continue
			}
			connection.Write([]byte(msg))
		}
	}

}
