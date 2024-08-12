.DEFAULT_GOAL := help

set-go-env:
	$(eval GOBIN=$(PWD)/bin)
	$(eval PATH=$(GOBIN):$(shell go env GOPATH):$(PATH))

local/install-tools: set-go-env ## install tools on local
	@./tools/install_tools.sh "./tools/tools.go"

docker/start-db: ## start db on docker
	docker compose up -d stlatica_db

local/create-migration-file: set-go-env ## create migration file on local
	$(GOBIN)/goose -dir=db/migrations create new_migration sql

local/migration-up: set-go-env ## run migration on local
	$(GOBIN)/goose -dir=./db/migrations mysql "root@tcp(localhost:4000)/stlatica" up;

docker/migration-up: ## run migration on docker
	docker compose run --rm goose bash -c "`goose -dir=./db/migrations mysql "root@tcp(127.0.0.1:4000)/stlatica" up`"

local/migration-down: set-go-env ## run migration down on local
	$(GOBIN)/goose -dir=./db/migrations mysql "root@tcp(localhost:4000)/stlatica" down	

docker/migration-down: ## run migration down on docker
	docker compose run --rm goose bash -c "`goose -dir=./db/migrations mysql "root@tcp(127.0.0.1:4000)/stlatica" down`"

local/generate-code-from-db: set-go-env ## generate code from db on local
	@./db/sqlboiler/generate_codes.sh

local/codegen-from-oapi: set-go-env ## generate code from openapi on local
	mkdir ../shared/openapi/tmp
	docker compose up redocly
	$(GOBIN)/oapi-codegen -generate server -package openapi ../shared/openapi/tmp/internalapi/openapi.yaml > app/controllers/internalapi/v1/openapi/internalapi_server.gen.go
	$(GOBIN)/oapi-codegen -generate types -package openapi ../shared/openapi/tmp/internalapi/openapi.yaml > app/controllers/internalapi/v1/openapi/internalapi_types.gen.go
	$(GOBIN)/oapi-codegen -generate spec -package openapi ../shared/openapi/tmp/internalapi/openapi.yaml > app/controllers/internalapi/v1/openapi/internalapi_spec.gen.go
	rm -r ../shared/openapi/tmp

local/lint: set-go-env ## lint on local
	$(GOBIN)/golangci-lint run ./...

local/start-internalapi: set-go-env ## start internalapi on local
	GO_ENV=local go run app/cmd/internalapi/main.go

docker/start-internalapi: ## start internalapi on docker
	SOURCE_DIR=./app/cmd/internalapi docker compose up -d stlatica_server

docker/stop-internalapi: set-go-env ## stop internalapi on docker
	SOURCE_DIR=./app/cmd/internalapi docker compose down stlatica_server

docker/start-object-storage: ## start object storage on docker
	docker compose up -d minio

docker/stop-object-storage: ## stop object storage on docker
	docker compose down minio

docker/start-mock-internalapi: ## start mock internalapi on docker
	docker run --rm -p 127.0.0.1:8080:8080 --mount type=bind,source="./",target="/go/src" -w /go/src -it golang:1.20.6 /bin/sh -c "go run mock/cmd/internalapi/main.go"

help: ## print this message ## make
	@printf "\033[36m%-30s\033[0m %-50s %s\n" "[command]" "[Description]"
	@grep -E '^[/a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
