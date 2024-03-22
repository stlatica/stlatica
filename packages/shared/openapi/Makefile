DOCKER_RUN := docker compose exec -it redocly
REDOCLY_OPTION := --host 0.0.0.0

help:
	# ここにヘルプを書く

external-build: up
	${DOCKER_RUN} redocly build-docs ${REDOCLY_OPTION} externalapi/openapi.yaml --output=build/external.html
external-bundle: up
	${DOCKER_RUN} redocly bundle ${REDOCLY_OPTION} externalapi/openapi.yaml --output=externalapi/openapi-bundled.yaml
external-lint: up
	${DOCKER_RUN} redocly lint ${REDOCLY_OPTION} externalapi/openapi.yaml
external-preview: up
	${DOCKER_RUN} redocly preview-docs ${REDOCLY_OPTION} externalapi/openapi.yaml -p 8081
external-serve:
	${DOCKER_RUN} serve build/external.html -p 8091

internal-build: up
	${DOCKER_RUN} redocly build-docs ${REDOCLY_OPTION} internalapi/openapi.yaml --output=build/internal.html
internal-bundle: up
	${DOCKER_RUN} redocly bundle ${REDOCLY_OPTION} internalapi/openapi.yaml --output=internalapi/openapi-bundled.yaml
internal-lint: up
	${DOCKER_RUN} redocly lint ${REDOCLY_OPTION} internalapi/openapi.yaml
internal-preview: up
	${DOCKER_RUN} redocly preview-docs ${REDOCLY_OPTION} internalapi/openapi.yaml -p 8080
internal-serve:
	${DOCKER_RUN} serve build/internal.html

internal-mock: internal-bundle
	${DOCKER_RUN} npm prism mock ./internalapi/openapi-bundled.yaml


up:
	docker compose up -d redocly

build:
	docker compose build redocly

down:
	docker compose down --remove-orphans