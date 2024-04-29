#!/bin/bash

# Set GOOS and GOARCH for Linux
GOOS=linux GOARCH=amd64 go build -o dist/clear-badge-cache.sh

# Set GOOS and GOARCH for Windows
GOOS=windows GOARCH=amd64 go build -o dist/clear-badge-cache.exe
