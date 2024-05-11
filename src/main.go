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
  -m, --model TEXT          Large Language Model to use. Use 'ppl --model' to display available options.
                            Default: llama-3-sonar-small-32k-online
  -p, --pattern TEXT        Pattern used to drive AI behavior. Use 'ppl --pattern' to display available options.
                            Default: None.
  -t, --temperature NUMBER  The amount of randomness in the response, valued between 0 inclusive and 2 exclusive. 
                            Higher values are more random. Lower values more deterministic.
                            Default: 1
  -m, --max-tokens NUMBER   The maximum number of completion tokens returned by the API.
  -c, --citations           Returns citations in the response.
                            Default: false (flag absent).
  -o, --output TEXT         The output format you get the response.
                            Available options: text (default), yaml, json
  -v, --version             Print version information and exit.

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
	ModelFlag           string
	PatternFlag         string
	TemperatureFlag     int
	MaxTokensFlag       int
	ReturnCitationsFlag bool
	OutputFlag          string
	VersionFlag         bool
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
