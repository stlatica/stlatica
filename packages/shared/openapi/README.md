# OpenAPI Files

# setup

1. install nodejs

nvm, volta 等を利用してインストールがおすすめです

2. install pnpm

```
npm install -g pnpm
```

3. install packages

```
pnpm install
```

# Commands

※ すべて internal を external に置き換えても動作します

## pnpm internal:preview

プレビューサーバーを立ち上げます  
ホットリロードで動作する模様

<http://localhost:8080/>

## pnpm internal:bundle

openapi 定義を lint します

## pnpm internal:bundle

openapi ファイルを 1 つにバンドルします。

※ generate 系ツールに投げる際、ファイルひとつでないと動かない場合があるので必要となります。

## pnpm internal:build

HTML をビルドします

(開発中は preview で済むので、実質的に CI 用です)

## pnpm internal:serve

ビルドした HTML を表示します。

（デバッグ用）

## pnpm internal:mock

prism でモックサーバーを開始します
