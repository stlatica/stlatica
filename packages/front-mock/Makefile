# quick-dev: docker/build_node node_modules
# 	docker compose run --rm --service-ports api_mock /bin/sh -c "pnpm dev"
# # --service-ports でcomposeのポートを使える


dev: dev/up
	docker compose exec -it api_mock bash

ladle: dev/up
	docker compose exec -it api_mock pnpm ladle:dev

dev/up: docker/build_node
	docker compose up --build -d api_mock

# TODO: 権限のその場しのぎ対応 https://github.com/stlatica/stlatica/issues/354
dev/permission:
	sudo chmod 777 -R .

down:
	docker compose down


# 以下内部用コマンド

node_modules: package.json pnpm-lock.yaml
	docker run --rm --mount type=bind,source=./,target=/app stlatica_front_node pnpm install --frozen-lockfile


docker/build_node: ../frontend/docker/Dockerfile.node.dev
	docker build -t stlatica_front_node -f ../frontend/docker/Dockerfile.node.dev .
# touch docker/build_node
