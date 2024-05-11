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

func help() {
	flag.Usage = func() {
		h := "Command line interface for Perplexity.ai\n\n"

		h += "Usage:\n"
		h += "  ppl [OPTIONS] [PROMPT]\n\n"

		h += "Options:\n"
		h += "  -m, --model      Large Language Model to use. Available: default, llama3, gpt4turbo, mistral\n"
		h += "  -p, --pattern    Pattern used to drive AI behavior. Use 'ppl --pattern' to display available options.\n"
		h += "      --version    Print version information\n\n"

		h += "Exit Codes:\n"
		h += fmt.Sprintf("  %d\t%s\n", exitOK, "OK")
		h += fmt.Sprintf("  %d\t%s\n", exitRuntimeError, "Failed at runtime")
		h += fmt.Sprintf("  %d\t%s\n", exitInputError, "Invalid user input provided")
		h += "\n"

		h += "Examples:\n"
		h += "  ppl \"What is the meaning of life?\"\n"
		h += "  ppl --pattern=creative \"What is the meaning of life?\"\n"
		h += "  ppl --output yaml \"What is the meaning of life?\"\n"

		fmt.Fprintf(os.Stderr, h)
	}
}

func main() {
	help()
	flag.Parse()
}
