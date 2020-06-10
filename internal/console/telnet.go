package console

import (
	"fmt"
	"github.com/LeLuxNet/Shelly/internal/input"
	"github.com/LeLuxNet/Shelly/pkg/sessions"
	"net"
	"os"
)

const (
	host = "localhost"
)

func Telnet(port string) {
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
	input.ReaderInput(sessions.NewSession(conn, conn, nil, false))
	err := conn.Close()
	if err != nil {
		fmt.Println("Error closing telnet connection: " + err.Error())
	}
}
