package main

import (
	"main/Handle"
	"net"
)

func main() {

	dial, err := net.Dial("tcp", "localhost:5678")
	Handle.HandleError(err)
	defer dial.Close()

	data := []byte("Hello mapren!")
	_, err = dial.Write(data)
	Handle.HandleError(err)

}
