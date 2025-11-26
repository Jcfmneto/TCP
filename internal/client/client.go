package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Start() {
	conn, err := net.Dial("tcp", "0.0.0.0:3030")
	if err != nil {
		fmt.Printf("Error to connect to the Server \n")
	}
	scanner := bufio.NewScanner(os.Stdin)

	go func() {

		for {
			readBuff := make([]byte, 1024)
			n, err := conn.Read(readBuff)
			if err != nil {
				fmt.Println("error to read data")
				return
			}
			fmt.Println("Read -> ", string(readBuff[:n]))
		}
	}()

	for scanner.Scan() {
		formated := scanner.Text()
		conn.Write([]byte(formated + "\n"))
	}

}
