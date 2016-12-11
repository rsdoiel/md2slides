#
# Makefile to compile mkslides for Mac OS X, Linux, Windows 7
# as well as Raspberry Pi Zero, 1,2, and 3.
#

build:
	go build -o bin/mkslides cmds/mkslides/mkslides.go

test:
	go test

install:
	env GOBIN=$(HOME)/bin go install cmds/mkslides/mkslides.go

status:
	git status

save:
	git commit -am "quick save"
	git push origin master

website:
	./mk-website.bash

clean:
	if [ -d bin ]; then /bin/rm -fR bin; fi
	if [ -d dist ]; then /bin/rm -fR dist; fi
	if [ -f mkslides-release.zip ]; then /bin/rm mkslides-release.zip; fi

release:
	./mk-release.bash	
	
publish:
	./publish.bash
