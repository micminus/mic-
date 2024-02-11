@echo off
set GO111MODULE=off
go build src/main.go
main run "c:/Users/valen/Documents/test/eee/langage_en_cours/test/input.micm"
del main.Exe