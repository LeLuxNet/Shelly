package main

import (
	"flag"
	"github.com/LeLuxNet/Shelly/internal/console"
	"github.com/LeLuxNet/Shelly/pkg/initialize"
	"strconv"
)

func main() {
	telnetPort := flag.Int("telnet", 0, "Open a telnet/tcp port to connect to shelly")
	flag.Parse()

	initialize.Init()

	if *telnetPort != 0 {
		console.Telnet(strconv.Itoa(*telnetPort))
	} else {
		console.Local()
	}
}
