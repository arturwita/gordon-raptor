build:
	go build -o ./bin/main ./src/cmd

run:
	go run ./src/cmd/main.go

unit_tests:
	go test -v ./src/internal/tests/unit...

e2e_tests:
	go test -v ./src/internal/tests/e2e...

lint:
	goimports -w $(shell find . -name '*.go')
