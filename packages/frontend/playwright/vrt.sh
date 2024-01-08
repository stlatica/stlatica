sudo docker run -p 9323:9323 --mount type=bind,source="$(pwd)"/playwright,target=/app/playwright -it playwright-vrt pnpm $@
