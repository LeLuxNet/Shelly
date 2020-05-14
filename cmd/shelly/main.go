package main

import (
	"flag"
	"fmt"
	"github.com/LeLuxNet/Shelly/internal/console"
	"github.com/LeLuxNet/Shelly/pkg/initialize"
	"github.com/LeLuxNet/Shelly/pkg/sessions"
	"os"
	"strconv"
)

func main() {
	noColor := flag.Bool("no-colors", false, "Disable colors")
	telnetPort := flag.Int("telnet", 0, "Open a telnet/tcp port to connect to shelly")
	silent := flag.Bool("silent", false, "Launch shelly silent")
	inception := flag.Bool("inception", false, "Allow to run shelly inside of shelly")
	flag.Parse()

	if !*inception && os.Getenv(initialize.SHELLY_RUNNING_ENV) == "1" {
		fmt.Println("You are trying to run shelly inside of shelly.")
		fmt.Println("If you are sure this is what you want run shelly again with the -inception argument")
		os.Exit(1)
	}

	initialize.Init()

	sessions.NoColors = *noColor
	sessions.Silent = *silent
	if *telnetPort != 0 {
		console.Telnet(strconv.Itoa(*telnetPort))
	} else {
		console.Local()
	}
}
