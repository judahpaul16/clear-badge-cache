@echo off

REM Set GOOS and GOARCH for Windows
SET GOOS=windows
SET GOARCH=amd64
go build -o dist/clear-badge-cache-win.exe

REM Optional: Reset environment variables (not strictly necessary)
SET GOOS=
SET GOARCH=

REM Set GOOS and GOARCH for Linux
SET GOOS=linux
SET GOARCH=amd64
go build -o dist/clear-badge-cache-linux.sh

REM Optional: Reset environment variables (not strictly necessary)
SET GOOS=
SET GOARCH=

REM Set GOOS and GOARCH for macOS
SET GOOS=darwin
SET GOARCH=amd64
go build -o dist/clear-badge-cache-mac.sh
