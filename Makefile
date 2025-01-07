.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: swagger_reset
swagger_reset:
	swag init

.PHONY: create_migraion
create_migration:
	migrate create -ext sql -dir migrations <name>

.PHONY: upload_migraion
upload_migration:
	migrate -path migrations -database "postgres://vendetta:vendetta@localhost:5432/vendetta_go?sslmode=disable" up