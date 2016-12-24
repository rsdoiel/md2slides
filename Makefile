#
# Makefile to compile mkslides for Mac OS X, Linux, Windows 7
# as well as Raspberry Pi Zero, 1,2, and 3.
#

PROJECT = mkslides

VERSION = $(shell grep -m1 'Version = ' $(PROJECT).go | cut -d\" -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

build:
	go build -o bin/mkslides cmds/mkslides/mkslides.go

test:
	go test

install:
	env GOBIN=$(HOME)/bin go install cmds/mkslides/mkslides.go

status:
	git status

save:
	git commit -am "Quick Save"
	git push origin $(BRANCH)

website:
	./mk-website.bash

clean:
	if [ -d bin ]; then /bin/rm -fR bin; fi
	if [ -d dist ]; then /bin/rm -fR dist; fi
	if [ -f mkslides-$(VERSION)-release.zip ]; then /bin/rm mkslides-$(VERSION)-release.zip; fi

release:
	./mk-release.bash	
	
publish:
	./publish.bash
