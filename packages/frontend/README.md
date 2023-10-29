# develop with Docker

```
docker compose up -d
docker exec -it frontend bash
pn install && pn dev
```

※ pnpm を pn にエイリアス張ってあります

# env

| tool     | version |
| -------- | ------- |
| volta ※1 | 1.1.1   |
| node     | 18.17.1 |
| pnpm     | 8.6.12  |

※ nvm等の他ツールがお好みであればnvm等でも可です

↑ 最終的にDocker化したら何もかも不要になりそう

# installs

## volta

nodejsの管理ツールです

https://docs.volta.sh/guide/getting-started

or

> winget install Volta

## node

nodejsそのままをインストールするとバージョン切り替えが辛いので、volta や nvm 経由で導入するのが基本になります

### with volta

自動で切り替わるのでインストール等不要です

### with nvm

> nvm install node@18.17.1

## pnpm

nodeのインストール後に実行します

> npm install -g pnpm@8.6.12

---

# quick start

> pnpm install
> pnpm dev

[http://localhost:3000](http://localhost:3000) にサーバーが立ちます

---

---

# TEST

単体テスト

> pnpm test

単発でテストを起動します

> pnpm test:snap

スナップショットを更新します

> pnpm test:watch

ウォッチモードでテストを起動します

> pnpm test:coverage

テストカバレッジを生成します

coverage/index.html に生成されるファイルを直接ブラウザで開くと確認できます

---

# plop

テンプレート生成ツール

> pnpm plop

起動します　あとは指示に従って入力したり選択するだけです

plop/plopfile.js に設定があります

# lint

推奨拡張を導入していれば vscode の保存時に自動整形されます。

手動でかける場合は以下で修正できます:

```
pnpm lint:fix
```
