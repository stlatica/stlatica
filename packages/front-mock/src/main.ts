import express from "express";
const app = express();
const port = 4010;
import type * as MyTypes from "./schema";
import { randomUUID } from "node:crypto";
import * as D from "date-fns";

app.get("/", (req, res) => {
  res.send("Hello World!");
});

app.get("/internal/v1/timelines/*", (req, res) => {
  type Res =
    MyTypes.paths["/internal/v1/timelines/{timeline_id}"]["get"]["responses"]["200"]["content"]["application/json"];

  const ResGen = (timeSub: number): Res[number] => {
    const time = D.subMilliseconds(new Date(), timeSub).toISOString();

    return {
      plat_id: randomUUID(),
      content: `メッセージ ${time}`,
      images: [],
      created_at: time,
    };
  };

  const r: Res = new Array(Math.floor(Math.random() * 10)).fill(0).map((_, i) => ResGen(i));

  res.send(r);
});

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});
