# argを受け取って実行したいのでわざとdocker-composeに書いていない
docker run -p 127.0.0.1:9323:9323 --mount type=bind,source="./",target=/app -it playwright-vrt /bin/sh -c "pnpm install --frozen-lockfile && pnpm $@"
