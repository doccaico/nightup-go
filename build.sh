#!/bin/sh

GOOS=linux GOARCH=amd64 go build -o nightup_linux -ldflags="-s -w" -trimpath ./cmd/nightup-go 

GOOS=windows GOARCH=amd64 go build -o nightup_windows.exe -ldflags="-s -w" -trimpath ./cmd/nightup-go 
