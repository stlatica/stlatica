cd ../shared/openapi
pnpm install --frozen-lockfile
pnpm internal:bundle

cd ../../frontend/
pnpm orval
