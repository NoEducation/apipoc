build:
	@go build -o bin/apipoc

run: build
	@./bin/apipoc

test:
	@go test -v ./..