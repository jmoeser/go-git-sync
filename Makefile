SHELL := /bin/sh

.PHONY : test

test:
	go test ./... -v
