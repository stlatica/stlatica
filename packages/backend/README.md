# Setup

## Go

Go の version 管理には`goenv`を使用するため、[こちら](https://github.com/go-nv/goenv/blob/master/INSTALL.md)の手順に従って`goenv`のインストールを行なってください。

`goenv`のインストールが完了したら、`goenv install`を実行することで`.go-version`に記載のある version の Go がインストールされます。

## DB

開発で使用する DB は docker コンテナ上に建てます。 \
以下のコマンドで利用可能です。

`make start-local-db`

# Structure

backend package ではクリーンアーキテクチャを採用しています。 \
クリーンアーキテクチャの詳細は[こちら](https://miro.com/app/board/uXjVNUb8IMw=/?share_link_id=136453257125)を参照してください。

package 構成は以下の通りです。

```
├── app
│ ├── adapters
│ ├── cmd // entrypoint
│ ├── controllers
│ ├── domains
│ │ ├── types
│ │ └── entities
│ ├── repositories
│ │ ├── dao
│ │ └── entities
│ └── usecases
└── その他(db, tools etc.)
```

# Tools

## goose

`goose`は Go 製のマイグレーションツールです。 \
以下のコマンドが使用可能です。

- `make create-migration-file`: db/migrations 配下に新しく migration ファイルを作成
- `make migration-up`: migration の version を上げる
- `make migration-down`: migration の version を下げる

## sqlboiler

Go 製の ORM です。 \
テーブル定義からコードの自動生成を行うこともできます。

## oapi-codegen

openapi の定義から Go のコードを自動生成します。 \
以下のコマンドで自動生成を行います。

`make codegen-from-oapi`

生成されるコードで使用するフレームワークはデフォルトの echo とします。

## Air

ホットリロードに対応した Web サーバを立てる Go 製のツールです。\
以下のコマンドで、Web サーバを立てることができます。

`make start-internalapi`

# 各種コマンドの実行

開発中に仕様するスクリプトやコマンドは`Makefile`に集約しています。\
新しくシェルスクリプトなど追加する場合、`Makefile`からも呼び出すことができる形にしてください。

