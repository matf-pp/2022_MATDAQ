.PHONY: clean build reqres run-reqres

all: build

bin/reqres: cmd/reqres/main.go
	go build -o $@ $<

build: bin/reqres

reqrep: bin/reqres
	$<

run-reqres:
	bin/reqres

clean: 
	go clean
	rm -f bin/reqres
