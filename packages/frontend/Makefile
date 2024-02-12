DOCKER_VRT_RUN := docker compose run --rm --service-ports playwright_vrt /bin/sh -c

quick-dev: docker/build_node node_modules
	docker compose run --rm --service-ports front_dev /bin/sh -c "pnpm dev"
# --service-ports でcomposeのポートを使える


dev: dev/up
	docker compose exec -it front_dev bash

ladle: dev/up
	docker compose exec -it front_dev pnpm ladle:dev

dev/up: docker/build_node node_modules
	docker compose up --build -d front_dev

down:
	docker compose down


docker/clear-cache:
	rm docker/build_node
	rm docker/build_playwright

docker/stop: 
	docker compose down --remove-orphans

vrt: docker/build_playwright
	${DOCKER_VRT_RUN} "pnpm install --frozen-lockfile --force && pnpm vrt"

vrt/update: docker/build_playwright
	${DOCKER_VRT_RUN} "pnpm install --frozen-lockfile --force && pnpm vrt:update"

# 以下内部用コマンド

node_modules: package.json pnpm-lock.yaml
	docker run --rm --mount type=bind,source=./,target=/app stlatica_front_node pnpm install --frozen-lockfile


docker/build_node: docker/Dockerfile.node.dev
	docker build -t stlatica_front_node -f docker/Dockerfile.node.dev .
# touch docker/build_node

docker/build_playwright: docker/Dockerfile.vrt
	docker build -t playwright-vrt --file docker/Dockerfile.vrt ./
# touch docker/build_playwright
