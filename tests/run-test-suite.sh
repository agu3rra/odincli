#!/bin/bash
# Run this from the repository's root folder.
set -e
go build -o odin ./...
go test -v ./... -coverpkg=./... -coverprofile=./tests/coverage.out 
go tool cover -html=./tests/coverage.out
