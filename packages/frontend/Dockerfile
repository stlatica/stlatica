FROM node:18.17.0-slim

WORKDIR /app

RUN npm install -g pnpm@8

RUN echo 'alias pn="pnpm"' >> ~/.bashrc
RUN echo 'alias px="pnpx"' >> ~/.bashrc

COPY package.json pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile
