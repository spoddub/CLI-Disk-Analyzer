build:
	go build -o ./bin/hexlet-path-size ./cmd/hexlet-path-size
tidy:
	go mod tidy
	go fmt ./...
	go vet ./...
test:
	go test ./...
lint:
	golangci-lint run
lint-fix:
	golangci-lint run --fix
