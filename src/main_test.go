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
				name: "no argument displays help",
				args: []string{
					"odin",
				},
				wantExitCode: common.ExitInputError,
				wantStdout:   "",
				wantStderr:   "Usage",
			},
			{
				name: "version display",
				args: []string{
					"odin",
					"version",
				},
				wantExitCode: common.ExitOK,
				wantStdout:   "odin version:",
				wantStderr:   "",
			},
			{
				name: "unexisting subcommand",
				args: []string{
					"odin",
					"foobar",
				},
				wantExitCode: common.ExitInputError,
				wantStdout:   "",
				wantStderr:   "Usage",
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
				readOut, writeOut, _ := os.Pipe()
				readErr, writeErr, _ := os.Pipe()
				os.Stdout = writeOut
				os.Stderr = writeErr

				// Run the main function or any specific function
				var gotExitCode int
				go func() {
					gotExitCode = run(testCase.args)
					writeOut.Close()
					writeErr.Close()
				}()

				// Read and check outputs
				outBytes, _ := io.ReadAll(readOut)
				errBytes, _ := io.ReadAll(readErr)

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
