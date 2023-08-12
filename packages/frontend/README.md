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

# jest

単体テスト

> pnpm test

単発でテストを起動します

> pnpm test:watch

ウォッチモードでテストを起動します

予期せぬデグレを防ぐため、開発中は常に立ち上げておくことを推奨します

> pnpm test:coverage

テストカバレッジを生成します

coverage/lconv-report/index.html に生成されるファイルを直接ブラウザ開くと確認できます

---

# plop

テンプレート生成ツール

> pnpm plop

起動します　あとは指示に従って入力したり選択するだけです

plop/plopfile.js に設定があります
