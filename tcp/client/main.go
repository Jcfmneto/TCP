package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	conn, err := net.Dial("tcp", "0.0.0.0:3030")
	if err != nil {
		fmt.Printf("error, %s", err)
	}

	for scanner.Scan() {
		txt := scanner.Text()
		conn.Write([]byte(txt))
	}

}
