export GOBIN := $(PWD)/bin
GO_PATH:=$(shell go env GOPATH)
export PATH:=$(GOBIN):$(GO_PATH):$(PATH)

install-tools:
	@./tools/install_tools.sh "./tools/tools.go"

start-local-db:
	docker-compose up -d stlatica_db

create-migration-file:
	goose -dir=db/migrations create new_migration sql

migration-up:
	goose -dir=./db/migrations mysql "root@tcp(localhost:4000)/stlatica" up

migration-down:
	goose -dir=./db/migrations mysql "root@tcp(localhost:4000)/stlatica" down	

generate-code-from-db:
	sqlboiler --config=db/sqlboiler.toml mysql

codegen-from-oapi:
	oapi-codegen -generate server -package openapi ../shared/openapi/internalapi/openapi.yaml > controllers/internalapi/v1/openapi/internalapi_server.gen.go
	oapi-codegen -generate types -package openapi ../shared/openapi/internalapi/openapi.yaml > controllers/internalapi/v1/openapi/internalapi_types.gen.go
	oapi-codegen -generate spec -package openapi ../shared/openapi/internalapi/openapi.yaml > controllers/internalapi/v1/openapi/internalapi_spec.gen.go
