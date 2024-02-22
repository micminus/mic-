@echo off
set GO111MODULE=off
go build src/main.go
main run "C:\Users\valen\Documents\GitHub\mic-\test\input.micm"
del main.Exe