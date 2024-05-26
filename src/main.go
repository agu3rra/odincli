package main

import (
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
	os.Exit(common.ExitInputError)
}

// Runs the cli and returns the exit code for the OS
// We keep it separate from main to facilitate test evaluation while maintaining it all in the same process as the tests
func run(args []string) int {
	if len(args) < 2 {
		help()
	}

	subcommand := args[1]
	switch subcommand {
	case "version":
		subcommands.Version()
		return common.ExitOK
	default:
		help()
	}
	return common.ExitOK
}

func main() {
	os.Exit(run(os.Args))
}
