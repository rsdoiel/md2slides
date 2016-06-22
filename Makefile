#
# Makefile to compile md2slides for Mac OS X, Linux, Windows 7
# as well as R-pi.
#

build:
	go build -v -o bin/md2slides cmds/md2slides/md2slides.go

install:
	env GOBIN=$(HOME)/bin go install -v cmds/md2slides/md2slides.go

clean:
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi

release:
	./mk-release.sh	
	

