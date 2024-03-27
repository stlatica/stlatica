import react from "@vitejs/plugin-react";
import tsconfigPaths from "vite-tsconfig-paths";
import { defineConfig } from "vitest/config";

export default defineConfig({
  plugins: [react(), tsconfigPaths()],
  test: {
    // ...
    environment: "jsdom",
    include: ["src/**/*.test.{js,ts,jsx,tsx}"],
  },
});
