# preview setup

1. install nodejs
2. install pnpm
3. install packages ※1

※1

```
pnpm install
```

# Commands

## pnpm internal:preview

プレビューサーバーを立ち上げます

<http://localhost:8080/>

ビルド自体はホットリロードされますが、ブラウザ側は手動更新が必要です。

## pnpm internal:bundle

openapi 定義を lint します

## pnpm internal:bundle

openapi ファイルを 1 つにバンドルします。

## pnpm internal:build

HTML をビルドします

(実質的に CI 用)

## pnpm internal:serve

ビルドした HTML を表示します
