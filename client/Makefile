.PHONY: clean build rc pd run-rc run-pd

all: build

bin/request-creator: cmd/request-creator/main.go
	go build -o $@ $<

bin/price-display: cmd/price-display/main.go
	go build -o $@ $<

bin/request-creator-server: cmd/request-creator-server/main.go
	go build -o $@ $<

build: bin/request-creator bin/price-display bin/request-creator-server

rc: bin/request-creator
	$<

pd: bin/price-display
	$<

rs: bin/request-creator-server
	$<

run-rc:
	bin/request-creator

run-pd:
	bin/price-display

run-rs:
	bin/request-creator-server

clean: 
	go clean
	rm -f bin/request-creator
	rm -f bin/price-display
	rm -f bin/request-creator-server
