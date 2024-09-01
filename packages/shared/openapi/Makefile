DOCKER_RUN := docker compose run --build --rm

.DEFAULT_GOAL := help


internal/preview: ## API定義確認サーバーを起動します(hot reload)
	docker compose up --build -d internal-preview
	# "http://localhost:2940"
external/preview: ## API定義確認サーバーを起動します(hot reload)
	docker compose up --build -d external-preview
	# "http://localhost:2941"
stop: ## API定義確認サーバーを停止します
	docker compose down

internal/lint: ## lintチェックします
	# ${DOCKER_RUN} redocly-cli redocly lint internalapi/openapi.yaml
external/lint: ## lintチェックします
	${DOCKER_RUN} redocly-cli redocly lint externalapi/openapi.yaml

internal/bundle: ## 1ファイルにバンドルします
	${DOCKER_RUN} redocly-cli redocly bundle internalapi/openapi.yaml --output=build/openapi-bundled-internal.yaml
external/bundle: ## 1ファイルにバンドルします
	${DOCKER_RUN} redocly-cli redocly bundle externalapi/openapi.yaml --output=build/openapi-bundled-external.yaml

internal/build: ## 定義をhtmlのドキュメントで出力します
	${DOCKER_RUN} redocly-cli redocly build-docs internalapi/openapi.yaml --output=build/internal.html
external/build: ## 定義をhtmlのドキュメントで出力します
	${DOCKER_RUN} redocly-cli redocly build-docs externalapi/openapi.yaml --output=build/external.html

internal/mock: internal/bundle ## prismでモックサーバーを起動します
	${DOCKER_RUN} -it -p 127.0.0.1:4010:4010 redocly-cli prism mock ./build/openapi-bundled-internal.yaml --host 0.0.0.0
external/mock: external/bundle ## prismでモックサーバーを起動します
	${DOCKER_RUN} -it -p 127.0.0.1:4010:4010 redocly-cli prism mock ./build/openapi-bundled-external.yaml --host 0.0.0.0

full-check: internal/build internal/bundle external/build external/bundle ## 全チェックを実施します

# https://ktrysmt.github.io/blog/write-useful-help-command-by-shell/
help: ## ヘルプを表示します
	@printf "\033[36m%-25s\033[0m %-50s %s\n" "[command]" "[Description]" 
	@grep -E '^[/a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}'
