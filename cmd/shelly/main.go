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

const Version = "0.1.0"

func main() {
	parser := argparse.NewParser("shelly", "The cross-platform shell")

	telnetPort := parser.Int("t", "telnet", &argparse.Options{Help: "Open a telnet/tcp port to connect to shelly"})

	noColor := parser.Flag("", "no-colors", &argparse.Options{Help: "Disable colors"})
	silent := parser.Flag("s", "silent", &argparse.Options{Help: "Launch shelly silent"})
	inception := parser.Flag("", "inception", &argparse.Options{Help: "Allow to run shelly inside of shelly"})
	errOverride := parser.Flag("", "err-override", &argparse.Options{Help: "Override the stderr channel with stdout"})

	version := parser.Flag("v", "version", &argparse.Options{Help: "Display the current version"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	if *version {
		fmt.Println(Version)
		return
	} else if !*inception && os.Getenv(initialize.RunningEnv) != "" {
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
		if *errOverride {
			console.Local(os.Stdin, os.Stdout, os.Stdout)
		} else {
			console.Local(os.Stdin, os.Stdout, os.Stderr)
		}
	}
}
