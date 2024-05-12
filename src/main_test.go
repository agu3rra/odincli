package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestVersionFlag(test *testing.T) {
	tests := []struct {
		name     string
		flag     string
		want     string
		exitCode int
	}{
		{"TestVersionLongFlag", "--version", "pplcli version:", 0},
		{"TestVersionShortFlag", "-v", "pplcli version:", 0},
	}

	for _, testCase := range tests {
		test.Run(testCase.name, func(test *testing.T) {
			cmd := exec.Command("./ppl", testCase.flag)
			output, err := cmd.CombinedOutput()
			if err != nil {
				if exiterr, ok := err.(*exec.ExitError); ok {
					if exiterr.ExitCode() != testCase.exitCode {
						test.Errorf(
							"Expected exit code %d, got %d", testCase.exitCode, exiterr.ExitCode())
					}
				} else {
					test.Fatalf("cmd.Run() failed with %s", err)
				}
			}

			if !strings.Contains(string(output), testCase.want) {
				test.Errorf("Unexpected output: %s, want %s", string(output), testCase.want)
			}
		})
	}
}
