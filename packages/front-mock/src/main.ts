import express from "express";
const app = express();
const port = 3000;
import type * as MyTypes from "./schema";

// app.use(
//   OpenApiValidator.middleware({
//     apiSpec: "../openapi/internalapi/openapi-bundled.yaml",
//     // apiSpec: "../openapi/internalapi/openapi.yaml",
//     validateRequests: true, // (default)
//     validateResponses: true, // false by default
//   }),
// );

app.get("/", (req, res) => {
  res.send("Hello World!");
});

app.get("/internal/v1/users/hoge", (req, res) => {
  type T = MyTypes.external["schemas/User.yaml"];
  const ss: T = {};

  res.send("Hello World!");
});

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});
