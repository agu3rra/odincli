package main

import (
	"flag"
	"fmt"
	"os"
)

// Exit codes
const (
	exitOK = iota
	exitRuntimeError
	exitInputError
)

// The help menu...
func help() {
	helpMessage := `Command line interface for Perplexity.ai

Usage:
  ppl [OPTIONS] [PROMPT]

Options:
  -m, --model      Large Language Model to use. Available: default, llama3, gpt4turbo, mistral
  -p, --pattern    Pattern used to drive AI behavior. Use 'ppl --pattern' to display available options.
      --version    Print version information

Exit Codes:
  0	OK
  1	Failed at runtime
  2	Invalid user input provided

Examples:
  ppl "What is the meaning of life?"
  ppl --pattern=creative "What is the meaning of life?"
  ppl --output yaml "What is the meaning of life?"
`

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, helpMessage)
	}
}

// CLI's input configuration
type Config struct {
	ModelFlag   string
	PatternFlag string
	VersionFlag bool
}

func main() {
	help()
	cfg := Config{}
	flag.StringVar(&cfg.ModelFlag, "model", "gpt4turbo", "")
	flag.StringVar(&cfg.PatternFlag, "pattern", "None", "")
	flag.BoolVar(&cfg.VersionFlag, "version", false, "")
	flag.Parse()
	os.Exit(exitOK)
}
