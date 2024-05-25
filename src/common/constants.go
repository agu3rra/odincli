package common

// Defaults
const (
	DefaultProvider    = "perplexity"
	DefaultModel       = "gpt4-omni"
	DefaultTemperature = 1.0
)

// Exit codes
const (
	ExitOK = iota
	ExitRuntimeError
	ExitInputError
)
