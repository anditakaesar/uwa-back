migrationpath = migrations

run-dev:
	@echo "${PWD}"
	@go run .

build:
	@go build

pre-commit:
	@gofmt -w ./..
	@go vet
	@go mod tidy

check-coverage:
	@go test ./...  -coverpkg=./... -coverprofile ./coverage.out
	@go tool cover -func ./coverage.out

mockery-all:
	@mockery --name=LoggerInterface --dir=internal/log

create-migration: ## Create new migration file. It takes parameter `file` as filename. Usage: `make create-migration file=add_column_time`
	@migrate create -ext sql -dir ${migrationpath} -seq $(file)

tidy-vendor:
	@go mod tidy
	@go mod vendor

