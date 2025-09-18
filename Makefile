build:
	go build -o bin/app ./...

run:
	go run ./src/main.go

test:
	go test ./...

lint:
	goimports -w .
