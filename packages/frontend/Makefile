.DEFAULT_GOAL := help

# << main commands >>

start-prod: ## start production server
	docker compose run --build -p 127.0.0.1:3000:3000 prod

dev: node_modules ## run remix server
	pnpm run dev

ladle: node_modules ## run ladle (storybook server)
	pnpm run ladle dev

lint: node_modules ## lint & format
	pnpm run lint

test: node_modules ## run vitest ui
	pnpm run test

build: node_modules ## run build
	pnpm run build

start: ## start builded server
	pnpm run start


vrt: node_modules ## run visual regression test
	pnpm run vrt

vrt/update: node_modules ## update visual regression snapshots
	pnpm run vrt:update

# << ci commands >>

# 実行時間順で制御できるように分割している
util/ci-checks: ci/lint ci/test ci/build ci/vrt ## run all ci checks

ci/lint: node_modules
	pnpm run lint:ci

ci/test: node_modules
	pnpm run test:ci

ci/vrt: vrt

ci/build: build

docker/vrt: ## 暫定措置: run 'pnpm vrt' in docker
	docker compose run --build --rm vrt bash -c "pnpm vrt"

docker/vrt-debug:
	docker compose run --build --rm vrt bash

docker/vrt-update:## 暫定措置: run 'pnpm vrt:update' in docker
	docker compose run --build --rm vrt bash -c "pnpm vrt:update"

# docker/vrt-build:
# 	docker compose build vrt

# << utility commands >>

help: ## print this message ## make
	@printf "\033[36m%-25s\033[0m %-50s %s\n" "[command]" "[Description]" 
	@grep -E '^[/a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}'


# << internal commands >>

node_modules: package.json pnpm-lock.yaml
	pnpm install --frozen-lockfile

