build:
	go build -o bin/app ./src/cmd

run:
	go run ./src/cmd/main.go

test:
	go test ./...

lint:
	goimports -w .
