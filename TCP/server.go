package main

import (
	"fmt"
	"main/Handle"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:1234")
	Handle.HandleError(err)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		Handle.HandleError(err)

		go handleClient(conn)
	}

}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	Handle.HandleError(err)

	fmt.Printf("Received: %s\n", buffer[:n])
}
