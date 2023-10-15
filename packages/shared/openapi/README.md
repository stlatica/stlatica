# preview setup

1. install nodejs
2. install pnpm
3. install packages ※1

※1

```
pnpm install
```

# Commands

## internal:preview

プレビューサーバーを立ち上げます

<http://localhost:8080/>

ビルド自体はホットリロードされますが、ブラウザ側は手動更新が必要です。

## internal:bundle

openapi ファイルを 1 つにバンドルします。

## internal:build

HTML をビルドします

(実質的に CI 用)

## internal:serve

ビルドした HTML を表示します
