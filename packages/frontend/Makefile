DOCKER_VRT_RUN := docker compose run --rm --service-ports playwright_vrt /bin/sh -c

.DEFAULT_GOAL := help

# << main commands >>

dev: util/build remix/node_modules ## run remix server ##
	docker compose run --rm -p 127.0.0.1:5173:5173 front_dev /bin/sh -c "pnpm run dev"

bash: util/build ## run bash (for package install etc.) ##
	docker compose run --rm -it front_dev bash

ladle: util/build remix/node_modules ## run ladle (storybook server)
	docker compose run --rm front_dev_ladle bash -c "pnpm run ladle dev"

ladle-no-hmr: util/build remix/node_modules ## ladle for docker windows / mac
	docker compose run --rm -p 127.0.0.1:61000:61000 -p 127.0.0.1:24678:24678 front_dev bash -c "pnpm run ladle dev"

lint: util/build remix/node_modules ## lint & format
	docker compose run --rm front_dev bash -c "pnpm run lint"

test: util/build remix/node_modules ## run vitest ui
	docker compose run --rm -p 127.0.0.1:51204:51204 front_dev /bin/sh -c "pnpm run test"

build: util/build remix/node_modules ## run build
	docker compose run --rm front_dev /bin/sh -c "pnpm run build"

start: util/build ## start builded server
	docker compose run --rm -p 127.0.0.1:3000:3000 front_dev /bin/sh -c "pnpm run start"


vrt: util/build remix/node_modules ## run visual regression test
	${DOCKER_VRT_RUN} "pnpm run vrt"

vrt/update: util/build remix/node_modules ## update visual regression snapshots
	${DOCKER_VRT_RUN} "pnpm run vrt:update"

# << ci commands >>

util/ci-checks: ci/lint ci/test ci/build ci/vrt ## run all ci checks

ci/lint: util/build remix/node_modules
	docker compose run --rm -p 127.0.0.1:51204:51204 front_dev /bin/sh -c "pnpm run lint:ci"

ci/test: util/build remix/node_modules
	docker compose run --rm -p 127.0.0.1:51204:51204 front_dev /bin/sh -c "pnpm run test:ci"

ci/vrt: vrt

ci/build: build

# << utility commands >>


# TODO: 権限のその場しのぎ対応 https://github.com/stlatica/stlatica/issues/354
util/permission: ## fix file permission (TODO #354)
	sudo chmod 777 -R .


# https://ktrysmt.github.io/blog/write-useful-help-command-by-shell/
help: ## print this message ## make
	@echo "Example operations by makefile."
	@echo ""
	@echo "Usage: make SUB_COMMAND argument_name=argument_value"
	@echo ""
	@echo "Command list:"
	@echo ""
	@printf "\033[36m%-30s\033[0m %-50s %s\n" "[Sub command]" "[Description]" "[Example]"
	@grep -E '^[/a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | perl -pe 's%^([/a-zA-Z_-]+):.*?(##)%$$1 $$2%' | awk -F " *?## *?" '{printf "\033[36m%-30s\033[0m %-50s %s\n", $$1, $$2, $$3}'

# docker/stop:
# 	docker compose down --remove-orphans

# << internal commands >>

util/build:
	docker compose build front_dev playwright_vrt

remix/node_modules: remix/package.json remix/pnpm-lock.yaml
	docker compose run --rm front_dev pnpm install --frozen-lockfile


