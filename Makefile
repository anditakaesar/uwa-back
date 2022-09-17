run-dev:
	@go run .

pre-commit:
	@gofmt -w ./..
	@go vet
	@go mod tidy