import { randomUUID } from "node:crypto";
import * as D from "date-fns";
import express from "express";
import type * as MyTypes from "./schema";
import { InitArray } from "./utils";

const app = express();
const port = 4010;

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});

// 動作確認用のルートアクセス設定
app.get("/", (req, res) => {
  res.send("Hello World!");
});

// URLを指定
app.get("/internal/v1/timelines/*", (req, res) => {
  // 型を取る
  type Res =
    MyTypes.paths["/internal/v1/timelines/{timeline_id}"]["get"]["responses"]["200"]["content"]["application/json"];

  // ジェネレータ関数を作成
  const ResGen = (timeSub: number): Res[number] => {
    const time = D.subMilliseconds(new Date(), timeSub).toISOString();

    return {
      plat_id: randomUUID(),
      content: `メッセージ ${time}`,
      images: [],
      created_at: time,
    };
  };

  // 戻り値を生成
  const r: Res = InitArray(10).map((_, i) => ResGen(i));

  // 返却
  res.send(r);
});
