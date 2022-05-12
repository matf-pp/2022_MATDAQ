.PHONY: clean build rc pd run-rc run-pd

REQUEST_CREATOR_SRC        = $(shell find ./internal/request-creator -type f -name '*.go')
# REQUEST_CREATOR_SERVER_SRC = $(shell find ./internal/request-creator-server -type f -name '*.go')
PRICE_DISPLAY_SRC          = $(shell find ./internal/price-display -type f -name '*.go')

all: build

bin/request-creator: cmd/request-creator/main.go $(REQUEST_CREATOR_SRC)
	go build -o $@ $<

bin/request-creator-server: cmd/request-creator-server/main.go $(REQUEST_CREATOR_SERVER_SRC)
	go build -o $@ $<

bin/price-display: cmd/price-display/main.go $(PRICE_DISPLAY_SRC)
	go build -o $@ $<

build: bin/request-creator bin/price-display bin/request-creator-server

rc: bin/request-creator
	$<

rs: bin/request-creator-server
	$<

pd: bin/price-display
	$<

run-rc:
	bin/request-creator

run-rs:
	bin/request-creator-server

run-pd:
	bin/price-display

clean: 
	go clean
	rm -f bin/request-creator
	rm -f bin/price-display
	rm -f bin/request-creator-server
