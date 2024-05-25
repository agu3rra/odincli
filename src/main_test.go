package main

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/agu3rra/odincli/src/common"
)

func TestHelpWhenAbsentSubcommand(test *testing.T) {
	test.Run("subcommands absent", func(test *testing.T) {
		cmd := exec.Command("./../odin")
		stdout, stderr := cmd.CombinedOutput()

		if stderr == nil {
			test.Fatal("when the help menu appears, it should be printed to stderr since it indicates an error")
		}

		wantMessage := "Usage"

		if !strings.Contains(string(stdout), wantMessage) {
			test.Errorf("Unexpected output: %s, want %s", string(stdout), wantMessage)
		}

		wantExitCode := common.ExitRuntimeError
		if exiterr, ok := stderr.(*exec.ExitError); ok {
			if exiterr.ExitCode() != wantExitCode {
				test.Errorf(
					"Expected exit code %d, got %d", wantExitCode, exiterr.ExitCode())
			}
		} else {
			test.Fatalf("cmd.Run() failed with %s", stderr)
		}
	})
}

func TestVersionSubcommand(test *testing.T) {
	tests := []struct {
		name         string
		subcommand   string
		wantMessage  string
		wantExitCode int
	}{
		{
			name:         "normal usage",
			subcommand:   "version",
			wantMessage:  "odin version:",
			wantExitCode: common.ExitOK,
		},
	}

	for _, testCase := range tests {
		test.Run(testCase.name, func(test *testing.T) {
			cmd := exec.Command("./../odin", testCase.subcommand)
			stdout, stderr := cmd.CombinedOutput()
			if stderr != nil {
				if exiterr, ok := stderr.(*exec.ExitError); ok {
					if exiterr.ExitCode() != testCase.wantExitCode {
						test.Errorf(
							"Expected exit code %d, got %d", testCase.wantExitCode, exiterr.ExitCode())
					}
				} else {
					test.Fatalf("cmd.Run() failed with %s", stderr)
				}
			}
			// we have an OK exit, so we check stdout has something we want to see
			if !strings.Contains(string(stdout), testCase.wantMessage) {
				test.Errorf("Unexpected output: %s, want %s", string(stdout), testCase.wantMessage)
			}
		})
	}
}
