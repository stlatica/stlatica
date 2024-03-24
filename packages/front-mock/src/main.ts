import express from "express";
const app = express();
const port = 4010;
import type * as MyTypes from "./schema";

app.get("/", (req, res) => {
  res.send("Hello World!");
});

app.get("/internal/v1/timelines/*", (req, res) => {
  type Res =
    MyTypes.paths["/internal/v1/timelines/{timeline_id}"]["get"]["responses"]["200"]["content"]["application/json"];

  const ResGen = (): Res[number] => {
    return {
      plat_id: "string",
      content: "",
      images: undefined,
      created_at: new Date().toISOString(),
    };
  };

  const r: Res = [ResGen()];

  res.send(r);
});

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});
