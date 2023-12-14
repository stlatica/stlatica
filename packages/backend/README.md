# Setup

## Go

Goのversion管理には`goenv`を使用するため、[こちら](https://github.com/go-nv/goenv/blob/master/INSTALL.md)の手順に従って`goenv`のインストールを行なってください。

`goenv`のインストールが完了したら、`goenv install`を実行することで`.go-version`に記載のあるversionのGoがインストールされます。

## DB

開発で使用するDBはdockerコンテナ上に建てます。 \
以下のコマンドで利用可能です。

`make start-local-db`

# Tools

## goose

`goose`はGo製のマイグレーションツールです。 \
以下のコマンドが使用可能です。

- `make create-migration-file`: db/migrations配下に新しくmigrationファイルを作成
- `make migration-up`: migrationのversionを上げる
- `make migration-down`: migrationのversionを下げる

## sqlboiler

Go製のORMです。 \
テーブル定義からコードの自動生成を行うこともできます。

## oapi-codegen

openapiの定義からGoのコードを自動生成します。 \
以下のコマンドで自動生成を行います。

`make codegen-from-oapi`

生成されるコードで使用するフレームワークはデフォルトのechoとします。


## Air
ホットリロードに対応したWebサーバを立てるGo製のツールです。\
以下のコマンドで、Webサーバを立てることができます。

`make start-internalapi`

# 各種コマンドの実行

開発中に仕様するスクリプトやコマンドは`Makefile`に集約しています。\
新しくシェルスクリプトなど追加する場合、`Makefile`からも呼び出すことができる形にしてください。
