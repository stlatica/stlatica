export GOBIN := $(PWD)/bin
GO_PATH:=$(shell go env GOPATH)
export PATH:=$(GOBIN):$(GO_PATH):$(PATH)

install-tools:
	@./tools/install_tools.sh "./tools/tools.go"

start-local-db:
	docker compose up -d stlatica_db

create-migration-file:
	$(GOBIN)/goose -dir=db/migrations create new_migration sql

migration-up:
	$(GOBIN)/goose -dir=./db/migrations mysql "root@tcp(localhost:4000)/stlatica" up

migration-down:
	$(GOBIN)/goose -dir=./db/migrations mysql "root@tcp(localhost:4000)/stlatica" down	

generate-code-from-db:
	@./db/sqlboiler/generate_codes.sh

codegen-from-oapi:
	$(GOBIN)/oapi-codegen -generate server -package openapi ../shared/openapi/internalapi/openapi.yaml > app/controllers/internalapi/v1/openapi/internalapi_server.gen.go
	$(GOBIN)/oapi-codegen -generate types -package openapi ../shared/openapi/internalapi/openapi.yaml > app/controllers/internalapi/v1/openapi/internalapi_types.gen.go
	$(GOBIN)/oapi-codegen -generate spec -package openapi ../shared/openapi/internalapi/openapi.yaml > app/controllers/internalapi/v1/openapi/internalapi_spec.gen.go
	
lint:
	$(GOBIN)/golangci-lint run ./...

start-internalapi:
	SOURCE_DIR=./app/cmd/internalapi docker compose up -d stlatica_server
