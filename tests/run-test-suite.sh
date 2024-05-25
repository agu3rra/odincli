#!/bin/bash
# Run this from the repository's root folder.
go build -o odin ./src
go test -v -coverprofile=./tests/coverage.out ./src
go tool cover -html=./tests/coverage.out
