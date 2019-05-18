.PHONY: all mod start

export GO111MODULE ?= on

all: mod start

mod:
	go mod download

start:
	go run app/api/main.go
