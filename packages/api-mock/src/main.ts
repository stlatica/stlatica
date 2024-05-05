import { randomUUID } from "node:crypto";
import * as D from "date-fns";
import express from "express";
import type * as MyTypes from "./schema";
import { InitArray } from "./utils";

const app = express();
const port = 4010;

const allowCrossDomain = function (req, res, next) {
  res.header("Access-Control-Allow-Origin", "*");
  res.header("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE");
  res.header(
    "Access-Control-Allow-Headers",
    "Content-Type, Authorization, access_token"
  );

  // intercept OPTIONS method
  if ("OPTIONS" === req.method) {
    res.send(200);
  } else {
    next();
  }
};
app.use(allowCrossDomain);

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});

// 動作確認用のルートアクセス設定
app.get("/", (req, res) => {
  res.send("Hello World!");
});

type Plat =
  MyTypes.paths["/internal/v1/plats/{plat_id}"]["get"]["responses"]["200"]["content"]["application/json"];

const history = new Map<string, Plat>();

const GenerateTimeline = () => {
  // 型を取る
  type Res =
    MyTypes.paths["/internal/v1/timelines"]["get"]["responses"]["200"]["content"]["application/json"];

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
  const r: Res = InitArray(Math.floor(Math.random() * 10)).map((_, i) =>
    ResGen(i)
  );

  r.forEach((x) => {
    history.set(x.plat_id, x);
  });

  return r;
};

// URLを指定
app.get("/internal/v1/timelines/", (req, res) => {
  const r = GenerateTimeline();
  // 返却
  res.send(r);
});

// URLを指定
app.get("/internal/v1/timelines/*", (req, res) => {
  const r = GenerateTimeline();
  // 返却
  res.send(r);
});

// URLを指定
app.get("/internal/v1/plats/:plat_id", (req, res) => {
  // 型を取る
  type Res =
    MyTypes.paths["/internal/v1/plats/{plat_id}"]["get"]["responses"]["200"]["content"]["application/json"];

  // 返却
  res.send(history.get(req.params.plat_id));
});

// http://localhost:4010/internal/v1/plats/2d13cbc6-de9c-439c-b4dd-c945dfc7628d
