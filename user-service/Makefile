.PHONY: clean build us run-us

all: build

bin/user-server: cmd/user-server/main.go
	go build -o $@ $<

build: bin/user-server

us: bin/user-server
	$<

run-us:
	bin/user-server

clean: 
	go clean
	rm -f bin/user-server
