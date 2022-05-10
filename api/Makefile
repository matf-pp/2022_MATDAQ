.PHONY: userService matchingEngine matchingEngineGo

all: build

build: userService matchingEngine

userService: ./user-service/userService.proto
	protoc $< \
		--go_out=. \
		--go_opt=paths=source_relative \
		--proto_path=.

matchingEngine: matchingEngineGo

matchingEngineGo: ./matching-engine/matchingEngine.proto
	protoc $< \
		--go_out=. \
		--go_opt=paths=source_relative \
		--proto_path=.
clean:
	go clean
