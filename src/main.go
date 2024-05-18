package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// Defaults
const (
	defaultModel       = "gpt4omni"
	defaultTemperature = 1.0
	defaultTokens      = 4096
)

// Exit codes
const (
	exitOK = iota
	exitRuntimeError
	exitInputError
)

var (
	logInfo = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logWarn = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	logErr  = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)

// The help menu...
func help() {
	helpMessage := `Command line interface for OpenAI's ChatGPT

Usage:
  gpterm [OPTIONS] [PROMPT]

Options:
  -m, --model TEXT          Large Language Model to use. Use 'ppl --model' to display available options.
                            Default: llama-3-sonar-small-32k-online
  -p, --pattern TEXT        Pattern used to drive AI behavior. Use 'ppl --pattern' to display available options.
                            Default: None.
  -t, --temperature NUMBER  The amount of randomness in the response, valued between 0 inclusive and 2 exclusive. 
                            Higher values are more random. Lower values more deterministic.
                            Default: 1
  -mt, --max-tokens NUMBER  The maximum number of completion tokens returned by the API.
                            Default: 4096
  -c, --citations           Returns citations in the response.
                            Default: true
  -o, --output TEXT         The output format you get the response.
                            Options: text, yaml, json
							Default: text
  -v, --version             Print version information and exit.

Exit Codes:
  0	OK
  1	Failed at runtime
  2	Invalid user input provided

Examples:
  gpterm "What is the meaning of life?"
  gpterm --pattern=write_essay "What is the meaning of life?"
  gpterm --output yaml "What is the meaning of life?"
`

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, helpMessage)
		os.Exit(exitRuntimeError)
	}
}

// CLI's input configuration
type Config struct {
	ModelFlag                string
	ModelFlagShort           string
	PatternFlag              string
	PatternFlagShort         string
	TemperatureFlag          float64
	TemperatureFlagShort     float64
	MaxTokensFlag            int
	MaxTokensFlagShort       int
	ReturnCitationsFlag      bool
	ReturnCitationsFlagShort bool
	OutputFlag               string
	OutputFlagShort          string
	VersionFlag              bool
	VersionFlagShort         bool
}

func main() {
	help()
	cfg := Config{}

	flag.StringVar(&cfg.ModelFlag, "model", defaultModel, "")
	flag.StringVar(&cfg.ModelFlagShort, "m", defaultModel, "")
	flag.StringVar(&cfg.PatternFlag, "pattern", "None", "")
	flag.StringVar(&cfg.PatternFlagShort, "p", "None", "")
	flag.Float64Var(&cfg.TemperatureFlag, "temperature", defaultTemperature, "")
	flag.Float64Var(&cfg.TemperatureFlagShort, "t", defaultTemperature, "")
	flag.IntVar(&cfg.MaxTokensFlag, "max-tokens", defaultTokens, "")
	flag.IntVar(&cfg.MaxTokensFlagShort, "mt", defaultTokens, "")
	flag.BoolVar(&cfg.ReturnCitationsFlag, "citations", true, "")
	flag.BoolVar(&cfg.ReturnCitationsFlagShort, "c", true, "")
	flag.StringVar(&cfg.OutputFlag, "output", "text", "")
	flag.StringVar(&cfg.OutputFlagShort, "o", "text", "")
	flag.BoolVar(&cfg.VersionFlag, "version", false, "")
	flag.BoolVar(&cfg.VersionFlagShort, "v", false, "")

	flag.Parse()

	if cfg.VersionFlag || cfg.VersionFlagShort {
		version, err := os.ReadFile("version")
		if err != nil {
			logErr.Printf("Error reading version information: %v\n", err)
			os.Exit(exitRuntimeError)
		}
		fmt.Printf("gpterm version: %s\n", string(version))
		os.Exit(exitOK)
	}
	os.Exit(exitOK)
}
