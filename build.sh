#!/usr/bin/env bash
RUN_NAME="github.oasis.ready"

mkdir -p output

go mod tidy
go build -o output/${RUN_NAME}