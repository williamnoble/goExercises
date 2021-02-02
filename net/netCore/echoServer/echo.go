package main

// curl -v telnet://127.0.0.1:1026
import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listen()
}

func listen() {
	listener, err := net.Listen("tcp", ":1026")
	if err != nil {
		fmt.Println("Failed to open on port 1026")
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting the connection")
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	// Create buffer for input (conn)
	reader := bufio.NewReader(conn)
	// Reader input to buffer til new line encountered
	data, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Println("Failed to read from socket")
		conn.Close()
	}
	response(data, conn)
}

func response(data []byte, conn net.Conn) {
	defer func() {
		conn.Close()
	}()
	conn.Write(data)
}
