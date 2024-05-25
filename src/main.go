package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/agu3rra/odincli/src/common"
	"github.com/agu3rra/odincli/src/subcommands"
)

// The help menu...
func help() {
	helpMessage := `ODIN: Command line interface for Artifical Intelligence

Usage:
  odin [SUBCOMMAND] [OPTIONS]

Subcommands:
  chat                      Default text based interaction with the AI.
  list                      Display available options for configurable flags.
  config                    Initial setup of OpenAI API key.
  purge                     Delete all chat history from ~/.odincli/history.
  version                   Display version info and exit.
Use "odin [SUBCOMMAND] --help" for more information 

Exit Codes:
  0	OK
  1	Failed at runtime
  2	Invalid user input provided
`
	fmt.Fprint(os.Stderr, helpMessage)
	os.Exit(common.ExitRuntimeError)
}

func main() {
	// Define subcommands
	// cmdChat := flag.NewFlagSet("chat", flag.ExitOnError)
	// cmdConfig := flag.NewFlagSet("config", flag.ExitOnError)
	// cmdList := flag.NewFlagSet("list", flag.ExitOnError)
	// cmdPurge := flag.NewFlagSet("purge", flag.ExitOnError)
	// cmdVersion := flag.NewFlagSet("version", flag.ExitOnError)

	if len(os.Args) < 2 {
		help()
	}

	subcommand := os.Args[1]
	switch subcommand {
	case "version":
		subcommands.Version()
	default:
		help()
	}
	flag.Parse()
}
