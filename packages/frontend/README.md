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

---

以下はデフォルトのREADMEそのまま

---

This is a [Next.js](https://nextjs.org/) project bootstrapped with [`create-next-app`](https://github.com/vercel/next.js/tree/canary/packages/create-next-app).

## Getting Started

First, run the development server:

```bash
npm run dev
# or
yarn dev
# or
pnpm dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

You can start editing the page by modifying `app/page.tsx`. The page auto-updates as you edit the file.

This project uses [`next/font`](https://nextjs.org/docs/basic-features/font-optimization) to automatically optimize and load Inter, a custom Google Font.

## Learn More

To learn more about Next.js, take a look at the following resources:

- [Next.js Documentation](https://nextjs.org/docs) - learn about Next.js features and API.
- [Learn Next.js](https://nextjs.org/learn) - an interactive Next.js tutorial.

You can check out [the Next.js GitHub repository](https://github.com/vercel/next.js/) - your feedback and contributions are welcome!
