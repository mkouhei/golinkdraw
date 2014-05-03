#!/usr/bin/make -f
# -*- makefile -*-

BIN := linkdraw
SRC := *.go
GOPKG := github.com/mkouhei/golinkdraw/
GOPATH := $(CURDIR)/_build:$(GOPATH)
export GOPATH


all: clean format test build

prebuild:
	install -d $(CURDIR)/_build/src/$(GOPKG)
	cp -a $(CURDIR)/*.go $(CURDIR)/_build/src/$(GOPKG)
	cp -a $(CURDIR)/modules $(CURDIR)/_build/src/$(GOPKG)


build: prebuild
	go build -o _build/$(BIN)


clean:
	@rm -rf _build


format:
	for src in $(SRC); do \
		gofmt $$src > $$src.tmp ;\
		mv -f $$src.tmp $$src ;\
	done


test: prebuild
	go test
