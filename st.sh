#!/bin/bash
clc -s
cat Version.dat
go mod tidy
go fmt .
staticcheck .
go vet .
golangci-lint run
echo -n "go test . "
go test .
git st
