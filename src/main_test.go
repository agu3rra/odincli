package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/agu3rra/odincli/src/common"
)

func TestCommandLine(test *testing.T) {

	test.Run("subcommands absent", func(test *testing.T) {
		tests := []struct {
			name         string
			args         []string
			wantExitCode int
			wantStdout   string
			wantStderr   string
		}{
			{
				name:         "no argument displays help",
				args:         []string{},
				wantExitCode: common.ExitInputError,
				wantStdout:   "",
				wantStderr:   "Usage",
			},
			{
				name: "version",
				args: []string{
					"version",
				},
				wantExitCode: common.ExitOK,
				wantStdout:   "odin version:",
				wantStderr:   "",
			},
		}
		for _, testCase := range tests {
			test.Run(testCase.name, func(test *testing.T) {
				// Backup real stdout and stderr
				oldStdout := os.Stdout
				oldStderr := os.Stderr
				defer func() {
					os.Stdout = oldStdout
					os.Stderr = oldStderr
				}()

				// Create a pipe to capture output
				rOut, wOut, _ := os.Pipe()
				rErr, wErr, _ := os.Pipe()
				os.Stdout = wOut
				os.Stderr = wErr

				// Run the main function or any specific function
				var gotExitCode int
				go func() {
					gotExitCode = run(testCase.args)
					wOut.Close()
					wErr.Close()
				}()

				// Read and check outputs
				outBytes, _ := io.ReadAll(rOut)
				errBytes, _ := io.ReadAll(rErr)

				// Convert expected output strings to byte slices
				wantStdoutBytes := []byte(testCase.wantStdout)
				wantStderrBytes := []byte(testCase.wantStderr)

				// Assert on the outputs
				if !bytes.Contains(outBytes, wantStdoutBytes) {
					test.Errorf("Unexpected stdout content: %s", string(outBytes))
				}
				if !bytes.Contains(errBytes, wantStderrBytes) {
					test.Errorf("Unexpected stderr content: %s", string(errBytes))
				}
				if gotExitCode != testCase.wantExitCode {
					test.Errorf("Unexpected exit code: %d; wanted: %d", gotExitCode, testCase.wantExitCode)
				}
			})
		}
	})
}
