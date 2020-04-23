package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
)

const (
	host = "localhost"
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
	telnetPort := flag.Int("telnet", 0, "Open a telnet/tcp port to connect to shelly")
	flag.Parse()

	initialize()

	if *telnetPort != 0 {
		telnet(strconv.Itoa(*telnetPort))
		return
	} else {
		localInput()
	}
}
