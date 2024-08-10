@echo off
go build -mod=vendor -ldflags -H=windowsgui cmd/game/game.go
