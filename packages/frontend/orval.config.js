module.exports = {
  api: {
    output: {
      mode: "tags-split",
      target: "src/openapi/",
      schemas: "src/openapi/model",
      client: "swr",
      mock: true,
      useTypeOverInterfaces: true,
      baseUrl: "http://localhost:4010",
    },
    input: {
      target: "../shared/openapi/internalapi/openapi-bundled.yaml",
    },
  },
};
