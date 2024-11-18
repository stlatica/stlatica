import { vitePlugin as remix } from "@remix-run/dev";
import { installGlobals } from "@remix-run/node";
import { vanillaExtractPlugin } from "@vanilla-extract/vite-plugin";
import { defineConfig } from "vite";
import tsconfigPaths from "vite-tsconfig-paths";

installGlobals();

export default defineConfig({
  plugins: [
    remix({
      ssr: false,
      appDirectory: "./src",
      ignoredRouteFiles: ["**/."],
    }),
    tsconfigPaths(),
    vanillaExtractPlugin(),
  ],
  server: {
    host: "0.0.0.0",
  },
});
