.PHONY: clean build rc pd run-rc run-pd

all: build

bin/request-creator: cmd/request-creator/main.go
	go build -o $@ $<

bin/price-display: cmd/price-display/main.go
	go build -o $@ $<

build: bin/request-creator bin/price-display

rc: bin/request-creator
	$<

pd: bin/price-display
	$<

run-rc:
	bin/request-creator

run-pd:
	bin/price-display

clean: 
	go clean
	rm -f bin/request-creator
	rm -f bin/price-display
