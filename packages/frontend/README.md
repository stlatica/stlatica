# quick start

```sh
pnpm install
pnpm dev
```

[http://localhost:3000](http://localhost:3000) にサーバーが立ちます

## quick start with Docker

```
docker compose up -d
docker exec -it frontend bash
pnpm install && pnpm dev
```

# env

| tool     | version |
| -------- | ------- |
| volta ※1 | 1.1.1   |
| node     | 18.17.1 |
| pnpm     | 8.6.12  |

※ nvm等の他ツールがお好みであればnvm等でも可です

# installs

## volta

nodejsの管理ツールです

https://docs.volta.sh/guide/getting-started

or

```sh
winget install Volta
```

## node

nodejsそのままをインストールするとバージョン切り替えが辛いので、volta や nvm 等node管理ツール経由で導入を推奨します。

### with volta

自動で切り替わるのでインストール等一切不要です

### with nvm

```sh
nvm install node@18.17.1
```

## pnpm

```sh
npm install -g pnpm@8.6.12
```

---

---

# Commands

## PR前チェック

CIで回る項目をローカルで検証します

```sh
# ci-checkの略
pnpm cc
```

## TEST

テストを起動します

vitest ui を採用しているためブラウザで結果確認等できます

```sh
pnpm test
```

## Visual Regression Testing

package/frontend 内で実行することを前提にスクリプトを作成しています

docker環境必須です

内容は1コマンドのため、中身を見て直接実行してもOKです（windows向け）

### 初回のみ

dockerイメージをビルドします。初回、あるいは更新があった際に再実行してください。

// make等で自動化したい

```sh
playwright/vrt-build.sh vrt
```

### 実行

```sh
playwright/vrt.sh vrt
```

### アップデート

```sh
playwright/vrt.sh vrt:update
```

---

## plop

テンプレートジェネレータです

plop/plopfile.js に設定があります。必要に応じて調整してください

```
pnpm plop
```

## lint

推奨拡張を導入していれば vscode の保存時に自動整形されます。

手動で整形をかける場合は以下で修正できます:

```
pnpm lint:fix
```

