package main

import (
	"fmt"
	"github.com/LeLuxNet/Shelly/internal/console"
	"github.com/LeLuxNet/Shelly/internal/initialize"
	"github.com/LeLuxNet/Shelly/pkg/sessions"
	"github.com/akamensky/argparse"
	"os"
	"strconv"
)

func main() {
	parser := argparse.NewParser("shelly.wasm", "The cross-platform shell")

	noColor := parser.Flag("", "no-colors", &argparse.Options{Help: "Disable colors"})
	telnetPort := parser.Int("t", "telnet", &argparse.Options{Help: "Open a telnet/tcp port to connect to shelly"})
	silent := parser.Flag("s", "silent", &argparse.Options{Help: "Launch shelly silent"})
	inception := parser.Flag("", "inception", &argparse.Options{Help: "Allow to run shelly inside of shelly"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	if !*inception && os.Getenv(initialize.RunningEnv) != "" {
		fmt.Println("You are trying to run shelly inside of shelly.")
		fmt.Println("If you are sure this is what you want run shelly again with the --inception argument")
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
