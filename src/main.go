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
  config                    Interactive setup process. Run at least once before using ODIN. Data saved to ~/.odin/config.
  purge                     Delete all chat history from ~/.odincli/history.
  version                   Display version info and exit.
Use "odin [SUBCOMMAND] --help" for more information 

Exit Codes:
  0	OK
  1	Failed at runtime
  2	Invalid user input provided
`
	fmt.Fprint(os.Stderr, helpMessage)
}

// Runs the cli and returns the exit code for the OS
// We keep it separate from main to facilitate test evaluation while maintaining it all in the same process as the tests
func run(args []string) int {
	if len(args) < 2 {
		help()
		return common.ExitInputError
	}

	subcommand := args[1]
	switch subcommand {
	case "version":
		subcommands.Version()
		return common.ExitOK
	default:
		help()
		return common.ExitInputError
	}
}

func main() {
	os.Exit(run(os.Args))
}
