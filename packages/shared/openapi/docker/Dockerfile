FROM node:22.12.0-slim

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"

RUN corepack enable
# RUN pnpm config set store-dir /home/node/.pnpm-store
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm add --global @redocly/cli@1.12.0 @stoplight/prism-cli@5.8.1
