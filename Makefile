run-dev:
	@go run .

pre-commit:
	@gofmt -w ./..
	@go vet
	@go mod tidy

check-coverage:
	@go test ./...  -coverpkg=./... -coverprofile ./coverage.out
	@go tool cover -func ./coverage.out