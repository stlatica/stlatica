FROM node:18.17.0-slim

WORKDIR /app

RUN npm install -g pnpm@8

RUN echo 'alias pn="pnpm"' >> ~/.bashrc
RUN echo 'alias px="pnpx"' >> ~/.bashrc


COPY . .
RUN pnpm install --frozen-lockfile

RUN node -v
RUN pnpm build
# RUN pnpm start
CMD ["pnpm", "start", "--port", "3010"]