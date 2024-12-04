import { vitePlugin as remix } from "@remix-run/dev";
import { installGlobals } from "@remix-run/node";
import { vanillaExtractPlugin } from "@vanilla-extract/vite-plugin";
import serverAdapter from "hono-remix-adapter/vite";
import { defineConfig } from "vite";
import tsconfigPaths from "vite-tsconfig-paths";

installGlobals();

export default defineConfig({
  plugins: [
    remix({
      appDirectory: "./src",
      ignoredRouteFiles: ["**/."],
    }),
    serverAdapter({
      entry: "./src/server/index.ts",
    }),
    tsconfigPaths(),
    vanillaExtractPlugin(),
  ],
  server: {
    host: "0.0.0.0",
  },
});
