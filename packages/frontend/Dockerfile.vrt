FROM node:20.10.0
WORKDIR /app

COPY package.json pnpm-lock.yaml ./
RUN npm install -g pnpm@8 && pnpm install --frozen-lockfile && pnpm exec playwright install-deps && pnpm exec playwright install

COPY ../ ./

EXPOSE 9323
EXPOSE 61000