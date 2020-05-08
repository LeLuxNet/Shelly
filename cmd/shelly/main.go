package main

import (
	"flag"
	"github.com/LeLuxNet/Shelly/internal/console"
	"github.com/LeLuxNet/Shelly/pkg/initialize"
	"github.com/LeLuxNet/Shelly/pkg/session"
	"strconv"
)

func main() {
	noColor := flag.Bool("no-colors", false, "Disable colors")
	telnetPort := flag.Int("telnet", 0, "Open a telnet/tcp port to connect to shelly")
	flag.Parse()

	initialize.Init()

	session.NoColors = *noColor
	if *telnetPort != 0 {
		console.Telnet(strconv.Itoa(*telnetPort))
	} else {
		console.Local()
	}
}
