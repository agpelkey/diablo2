build:
	go build -o bin/diablo2

run: build
	./bin/diablo2

test:
	@go test -v ./...