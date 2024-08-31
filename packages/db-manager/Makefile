.DEFAULT_GOAL := help

# << main commands >>

dev: node_modules ## run remix server
	pnpm run dev

lint: node_modules ## lint & format
	pnpm run lint
	pnpm run check
	pnpm run format

test: node_modules ## run vitest ui
	pnpm run test

build: node_modules ## run build
	pnpm run build

start: ## start builded server
	pnpm run start

# << utility commands >>

help: ## print this message ## make
	@printf "\033[36m%-25s\033[0m %-50s %s\n" "[command]" "[Description]" 
	@grep -E '^[/a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}'


# << internal commands >>

node_modules: package.json pnpm-lock.yaml
	pnpm install --frozen-lockfile

