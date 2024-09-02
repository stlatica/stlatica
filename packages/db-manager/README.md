# Stlatica DB Mangement Tool (SDMT)

## 起動

packages/backend のDBを先に起動してください。```make docker/start-db docker/migration-up```

> make docker/start

## 開発方法

ネットワークの都合上 devcontainer 前提です。

devcontainer使わず直接編集したい場合は環境変数を調整してください。

## favicon memo

https://favicon.io/favicon-generator/

```
#5AF

pacifico

110
```

## create-svelte

Everything you need to build a Svelte project, powered by [`create-svelte`](https://github.com/sveltejs/kit/tree/main/packages/create-svelte).

## Creating a project

If you're seeing this, you've probably already done this step. Congrats!

```bash
# create a new project in the current directory
npm create svelte@latest

# create a new project in my-app
npm create svelte@latest my-app
```

## Developing

Once you've created a project and installed dependencies with `npm install` (or `pnpm install` or `yarn`), start a development server:

```bash
npm run dev

# or start the server and open the app in a new browser tab
npm run dev -- --open
```

## Building

To create a production version of your app:

```bash
npm run build
```

You can preview the production build with `npm run preview`.

> To deploy your app, you may need to install an [adapter](https://kit.svelte.dev/docs/adapters) for your target environment.
