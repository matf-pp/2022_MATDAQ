.PHONY: clean build us run-us

all: build

USER_SERVICE_SRC = $(shell find ./internal/ -type f -name '*.go')
 
bin/user-server: cmd/user-server/main.go $(USER_SERVICE_SRC)
	go build -o $@ $<

build: bin/user-server

us: bin/user-server
	$<

run-us:
	bin/user-server

clean: 
	go clean
	rm -f bin/user-server
