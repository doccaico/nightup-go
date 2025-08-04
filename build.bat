@echo off

setlocal

set GOOS=windows
set GOARCH=amd64

go build -o nightup_windows.exe -ldflags="-s -w" -trimpath .\cmd\nightup-go 

set GOOS=linux
set GOARCH=amd64

go build -o nightup_linux -ldflags="-s -w" -trimpath .\cmd\nightup-go 
