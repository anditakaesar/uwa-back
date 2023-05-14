migrationpath = migrations

run-dev:
	@go run .

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
	@ls -x ${migrationpath}/*.sql | tail -1 | awk -F"${migrationpath}/" '{print $$2}' | awk -F"_" '{print $$1}' | { read cur_v; expr $$cur_v + 1; } | { read new_v; printf "%06d" $$new_v; } | { read v; touch ${migrationpath}/$$v"_$(file)".up.sql; touch ${migrationpath}/$$v"_$(file)".down.sql; }

tidy-vendor:
	@go mod tidy
	@go mod vendor

prod-build:
	@go build -o ../app
	@cp -r ./migrations/ ~/
