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
			stdout, stderr := cmd.CombinedOutput()
			if stderr != nil {
				if exiterr, ok := stderr.(*exec.ExitError); ok {
					if exiterr.ExitCode() != testCase.exitCode {
						test.Errorf(
							"Expected exit code %d, got %d", testCase.exitCode, exiterr.ExitCode())
					}
				} else {
					test.Fatalf("cmd.Run() failed with %s", stderr)
				}
			}
			// we have an OK exit, so we check stdout has something we want to see
			if !strings.Contains(string(stdout), testCase.want) {
				test.Errorf("Unexpected output: %s, want %s", string(stdout), testCase.want)
			}
		})
	}
}
