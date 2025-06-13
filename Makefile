.PHONY: build run-server run-client clean

build:
	go build -o bin/server cmd/server/main.go
	go build -o bin/client cmd/client/main.go

run-server: build
	./bin/server

run-client: build
	./bin/client

clean:
	rm -rf bin/ 