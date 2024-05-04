DOCKER_RUN := docker compose run -it redocly

.DEFAULT_GOAL := help

internal/preview: docker-build ## API定義確認サーバーを起動します
	docker compose run -it -p 127.0.0.1:8080:8080 redocly redocly preview-docs --host 0.0.0.0 internalapi/openapi.yaml -p 8080
external/preview: docker-build ## API定義確認サーバーを起動します
	docker compose run -it -p 127.0.0.1:8081:8081 redocly redocly preview-docs --host 0.0.0.0 externalapi/openapi.yaml -p 8081

internal/lint: docker-build ## lintチェックします
	docker compose run -it redocly redocly lint internalapi/openapi.yaml
external/lint: docker-build ## lintチェックします
	docker compose run -it redocly redocly lint externalapi/openapi.yaml

internal/bundle: docker-build ## 1ファイルにバンドルします
	docker compose run -it redocly redocly bundle internalapi/openapi.yaml --output=build/openapi-bundled-internal.yaml
external/bundle: docker-build ## 1ファイルにバンドルします
	docker compose run -it redocly redocly bundle externalapi/openapi.yaml --output=build/openapi-bundled-external.yaml

internal/build: docker-build ## 定義をhtmlのドキュメントで出力します
	docker compose run -it redocly redocly build-docs internalapi/openapi.yaml --output=build/internal.html
external/build: docker-build ## 定義をhtmlのドキュメントで出力します
	docker compose run -it redocly redocly build-docs externalapi/openapi.yaml --output=build/external.html

internal/mock: docker-build internal/bundle ## prismでモックサーバーを起動します
	docker compose run -it -p 127.0.0.1:4010:4010 redocly prism mock ./internalapi/openapi-bundled.yaml --host 0.0.0.0
external/mock: docker-build external/bundle ## prismでモックサーバーを起動します
	docker compose run -it -p 127.0.0.1:4010:4010 redocly prism mock ./externalapi/openapi-bundled.yaml --host 0.0.0.0

docker-build:
	docker compose build 


# https://ktrysmt.github.io/blog/write-useful-help-command-by-shell/
help: ## print this message
	@printf "\033[36m%-25s\033[0m %-50s %s\n" "[command]" "[Description]" 
	@grep -E '^[/a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}'