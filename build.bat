@echo off
rem set OLD_GO_PATH="%GOPATH%"
rem set OLD_GO_BIN="%GOBIN%"

set CURRENT_PATH=%CD%
set SRC_PATH="%CURRENT_PATH%"
set GO111MODULE=auto
set MAIN_GO="main.go"
rem set GOPATH="%CURRENT_PATH%"


echo formating...
gofmt -w "%SRC_PATH%"
echo building...
cd %SRC_PATH%
go build "%MAIN_GO%"
echo build done...
cd %CURRENT_PATH%

rem set GOPATH="%OLD_GO_PATH%"
rem set GOBIN="%OLD_GO_BIN%"