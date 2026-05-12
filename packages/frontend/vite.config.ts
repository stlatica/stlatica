import { vitePlugin as remix } from "@remix-run/dev";
import { installGlobals } from "@remix-run/node";
import { vanillaExtractPlugin } from "@vanilla-extract/vite-plugin";
import serverAdapter from "hono-remix-adapter/vite";
import { defineConfig } from "vite";
import tsconfigPaths from "vite-tsconfig-paths";

installGlobals();

const ignoredRouteFiles = [
  "**/.*",
  "**/*.test.js",
  "**/*.test.jsx",
  "**/*.test.ts",
  "**/*.test.tsx",
  "**/*.spec.js",
  "**/*.spec.jsx",
  "**/*.spec.ts",
  "**/*.spec.tsx",
];

// Keep Vite-owned module requests away from the Hono dev middleware.
const serverAdapterExclude = [
  /.*\.css(\?.*)?$/,
  /.*\.ts(\?.*)?$/,
  /.*\.tsx(\?.*)?$/,
  /^\/@.+$/,
  /\?t=\d+$/,
  /^\/favicon\.ico$/,
  /^\/static\/.+/,
  /^\/node_modules\/.*/,
  "/assets/**",
  "/app/**",
  "/src/app/**",
];

export default defineConfig({
  plugins: [
    remix({
      appDirectory: "./src",
      ignoredRouteFiles,
    }),
    serverAdapter({
      entry: "./src/server/index.ts",
      exclude: serverAdapterExclude,
    }),
    tsconfigPaths(),
    vanillaExtractPlugin(),
  ],
  server: {
    host: "0.0.0.0",
  },
});
