.PHONY: all build test clean

all: build test

build:
	  go build -o boatapp cmd/boatapp/main.go

test:
	  go test ./...

clean:
	  rm -f boatapp

