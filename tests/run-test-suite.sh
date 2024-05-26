#!/bin/bash
# Run this from the repository's root folder.
set -e
go test -v ./... -coverpkg=./... -coverprofile=./tests/coverage.out 
go tool cover -html=./tests/coverage.out -o ./tests/coverage.html
