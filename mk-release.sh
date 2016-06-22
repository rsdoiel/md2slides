#!/bin/bash
#
# Make releases for Linux/amd64, Linux/ARM7 (Raspberry Pi), Windows, and Mac OX X (darwin)
#
env GOOS=linux GOARCH=amd64 go build -v -o dist/linux_amd64/md2slides cmds/md2slides/md2slides.go
env GOOS=linux GOARCH=arm GOARM=7 go build -v -o dist/raspberrypi/md2slides cmds/md2slides/md2slides.go
env GOOS=windows GOARCH=amd64 go build -v -o dist/windows/md2slides.exe cmds/md2slides/md2slides.go
env GOOS=darwin	GOARCH=amd64 go build -v -o dist/maxosx/md2slides cmds/md2slides/md2slides.go

