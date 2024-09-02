.DEFAULT_GOAL := help

# << production command >>

docker/start: ## start docker container
	docker compose up -d --build --remove-orphans prod

docker/logs: ## show docker logs
	docker compose logs -f

docker/stop: ## stop docker container
	docker compose down

# << main commands >>

dev: node_modules ## run remix server
	pnpm run dev

lint: node_modules ## lint & format
	pnpm run format
	pnpm run lint
	pnpm run check

test: node_modules ## run vitest ui
	pnpm run test

build: node_modules ## run build
	pnpm run build

start: ## start builded server
	pnpm run start

prisma-sync: ## sync prisma schema
	pnpm prisma db pull
	pnpm prisma generate

# << utility commands >>

help: ## print this message ## make
	@printf "\033[36m%-25s\033[0m %-50s %s\n" "[command]" "[Description]" 
	@grep -E '^[/a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}'


# << internal commands >>

node_modules: package.json pnpm-lock.yaml
	pnpm install --frozen-lockfile

