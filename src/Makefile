.DEFAULT_GOAL := build

fmt:
	go fmt GoScrap.go
.PHONY:fmt

lint: fmt
	golint GoScrap.go
.PHONY:lint

vet: fmt
	go vet GoScrap.go
.PHONY:vet

build: vet
	go build GoScrap.go
.PHONY:build
