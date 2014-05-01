#!/usr/bin/make -f
# -*- makefile -*-

BIN := linkdraw
SRC := *.go

all:
	go build -o _build/$(BIN)


clean:
	@rm -rf _build


format:
	for src in $(SRC); do \
		gofmt $$src > $$src.tmp ;\
		mv -f $$src.tmp $$src ;\
	done
