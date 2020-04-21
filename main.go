package main

import (
	"fmt"
	"net"
	"os"
)

const (
	host = "localhost"
	port = "3300"
)

func telnet(port string) {
	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		os.Exit(1)
	}

	defer listener.Close()
	fmt.Println("Listening on " + host + ":" + port)

	for {
		conn, err := listener.Accept()
		fmt.Println("Accepted connection")
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	ReaderInput(NewInOutErr(conn, conn, nil, false))
}

func main() {
	// telnet(port)
	localInput()
}
