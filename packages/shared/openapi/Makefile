DOCKER_RUN := docker compose exec -it redocly

x: help


internal-preview: docker-up
	${DOCKER_RUN} redocly preview-docs --host 0.0.0.0 internalapi/openapi.yaml -p 8080
external-preview: docker-up
	${DOCKER_RUN} redocly preview-docs --host 0.0.0.0 externalapi/openapi.yaml -p 8081

internal-lint: docker-up
	${DOCKER_RUN} redocly lint internalapi/openapi.yaml
external-lint: docker-up
	${DOCKER_RUN} redocly lint externalapi/openapi.yaml

internal-bundle: docker-up
	${DOCKER_RUN} redocly bundle internalapi/openapi.yaml --output=internalapi/openapi-bundled.yaml
external-bundle: docker-up
	${DOCKER_RUN} redocly bundle externalapi/openapi.yaml --output=externalapi/openapi-bundled.yaml


internal-build: docker-up
	${DOCKER_RUN} redocly build-docs  internalapi/openapi.yaml --output=build/internal.html
external-build: docker-up
	${DOCKER_RUN} redocly build-docs externalapi/openapi.yaml --output=build/external.html

internal-serve:
	${DOCKER_RUN} serve build/internal.html --host 0.0.0.0 -p 8090
external-serve:
	${DOCKER_RUN} serve build/external.html --host 0.0.0.0 -p 8091


internal-mock: internal-bundle
	${DOCKER_RUN} prism mock ./internalapi/openapi-bundled.yaml
external-mock: internal-bundle
	${DOCKER_RUN} prism mock ./externalapi/openapi-bundled.yaml


docker-up:
	docker compose up -d redocly

docker-build:
	docker compose build redocly

docker-down:
	docker compose down --remove-orphans


help:
	# > internal-preview
	# > external-preview
	# API定義確認サーバーを起動します
	#
	# > internal-lint
	# > external-lint
	# lintチェックします
	#
	# > internal-bundle
	# > external-bundle
	# 1ファイルにバンドルします
	#
	# > internal-build
	# > external-build
	# 定義をhtmlのドキュメントで出力します
	#
	# > internal-serve
	# > external-serve
	# buildしたファイルを確認します
	#
	# > internal-mock
	# > external-mock
	# prismでモックサーバーを起動します