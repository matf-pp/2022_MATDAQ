.PHONY: clean build reqrep run-reqrep

all: build

bin/reqrep: cmd/reqrep
	go build -o $@ $<

build: bin/reqrep

reqrep: bin/reqrep
	$<

run-reqrep:
	bin/reqrep

clean: 
	go clean
	rm -f bin/reqres
