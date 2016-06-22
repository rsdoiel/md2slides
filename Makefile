#
# Makefile to compile md2slides for Mac OS X, Linux, Windows 7
# as well as R-pi.
#

build:
	mkdir -p bin
	go build -v -o bin/md2slides cmds/md2slides/md2slides.go

install:
	env GOBIN=$(HOME)/bin go install -v cmds/md2slides/md2slides.go

release:
	./mk-release.sh	
	

