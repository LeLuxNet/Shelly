package main

import (
	"flag"
	"github.com/LeLuxNet/Shelly/internal/console"
	"github.com/LeLuxNet/Shelly/pkg/initialize"
	"github.com/LeLuxNet/Shelly/pkg/sessions"
	"strconv"
)

func main() {
	noColor := flag.Bool("no-colors", false, "Disable colors")
	telnetPort := flag.Int("telnet", 0, "Open a telnet/tcp port to connect to shelly")
	silent := flag.Bool("silent", false, "Launch shelly silent")
	flag.Parse()

	initialize.Init()

	sessions.NoColors = *noColor
	sessions.Silent = *silent
	if *telnetPort != 0 {
		console.Telnet(strconv.Itoa(*telnetPort))
	} else {
		console.Local()
	}
}
