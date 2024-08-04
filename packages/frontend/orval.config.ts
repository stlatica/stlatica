import { defineConfig } from "orval";

export default defineConfig({
  api: {
    output: {
      // .msw を別で生成するため
      mode: "split",
      target: "src/openapi/",
      client: "swr",
      mock: true,
      // useTypeOverInterfaces: true,
      baseUrl: "http://localhost:4010",
    },
    input: {
      target: "../shared/openapi/build/openapi-bundled-internal.yaml",
    },
  },
});
