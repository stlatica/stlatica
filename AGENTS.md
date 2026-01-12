# AGENTS

このリポジトリは monorepo です。作業対象のパッケージ配下に合わせて README/Makefile を参照してください。

## リポジトリ構成

- `docs/`: 設計資料など
- `packages/`:
  - `backend/`: Go + MySQL のバックエンド
  - `frontend/`: フロントエンド（TypeScript）
  - `shared/openapi/`: OpenAPI 定義
  - `api-mock/`: フロント向けモックサーバ
  - `db-manager/`: Stlatica DB Management Tool (SDMT)

## 開発フロー（起動/テスト/ビルド）
- 全体ビルド（レビュー用）: `packages/README.md` の `docker compose build --no-cache && docker compose up -d`

- Backend (`packages/backend/Makefile`)
  - 起動: `make local/start-internalapi` / `make docker/start-internalapi`
  - DB: `make docker/start-db`（停止は `make docker/stop-db`）
  - Lint: `make local/lint`
  - マイグレーション: `make local/create-migration-file`, `make local/migration-up`, `make local/migration-down`
  - OpenAPI コード生成: `make local/codegen-from-oapi`
- Frontend (`packages/frontend/Makefile`)
  - 起動: `make dev`（prod は `make start-prod`）
  - テスト: `make test` / VRT は `make vrt`
  - Lint: `make lint`
  - ビルド: `make build`（起動は `make start`）
- API Mock (`packages/api-mock/Makefile`)
  - 起動: `make run` / `make dev`
  - Ladle: `make ladle`
- DB Manager (`packages/db-manager/Makefile`)
  - 起動: `make dev` / Docker は `make docker/start`
  - テスト: `make test`
  - Lint: `make lint`
  - ビルド: `make build`（起動は `make start`）
- OpenAPI (`packages/shared/openapi/Makefile`)
  - Preview: `make internal/preview` / `make external/preview`（停止は `make stop`）
  - Lint: `make external/lint`
  - ビルド: `make internal/build` / `make external/build`

## Coding & Architecture Guidelines

- golang
  - クリーンアーキテクチャを守る
  - usecase はアプリケーションロジックのみ、外部依存はinfrastructure に隔離する

## 作業メモ

- 追加のスクリプトは `packages/backend/Makefile` から呼べる形にまとめることを推奨
- 仕様/詳細は各パッケージの README を優先



