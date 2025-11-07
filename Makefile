.DEFAULT_GOAL := build

.PHONY:fmt vet built
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o todo

clean:
	go clean