export GOBIN := $(PWD)/bin
GO_PATH:=$(shell go env GOPATH)
export PATH:=$(GOBIN):$(GO_PATH):$(PATH)

list:
	@grep "^[a-zA-Z\-]*:" Makefile | grep -v "grep" | sed -e 's/^/make /' | sed -e 's/://'

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
	mkdir ../shared/openapi/tmp
	docker compose up redocly
	$(GOBIN)/oapi-codegen -generate server -package openapi ../shared/openapi/tmp/internalapi/openapi.yaml > app/controllers/internalapi/v1/openapi/internalapi_server.gen.go
	$(GOBIN)/oapi-codegen -generate types -package openapi ../shared/openapi/tmp/internalapi/openapi.yaml > app/controllers/internalapi/v1/openapi/internalapi_types.gen.go
	$(GOBIN)/oapi-codegen -generate spec -package openapi ../shared/openapi/tmp/internalapi/openapi.yaml > app/controllers/internalapi/v1/openapi/internalapi_spec.gen.go
	rm -r ../shared/openapi/tmp

lint:
	$(GOBIN)/golangci-lint run ./...

start-local-internalapi:
	GO_ENV=local go run app/cmd/internalapi/main.go

start-docker-internalapi:
	SOURCE_DIR=./app/cmd/internalapi docker compose up -d stlatica_server

stop-docker-internalapi:
	SOURCE_DIR=./app/cmd/internalapi docker compose down stlatica_server


start-object-storage:
	docker compose up minio

stop-object-storage:
	docker compose down minio

# モックサーバーの起動
start-mock-internalapi:
	docker run --rm -p 127.0.0.1:8080:8080 --mount type=bind,source="./",target="/go/src" -w /go/src -it golang:1.20.6 /bin/sh -c "go run mock/cmd/internalapi/main.go"

