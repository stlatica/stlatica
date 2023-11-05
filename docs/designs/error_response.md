# エラー設計

APIのresponseで返すエラーの設計について記載する。

## Format
エラーコードとエラーの詳細を記載したメッセージを返す。
エラーコードについては[Error Code](#error-code)を参照。

```
{
  "code": "ERROR_CODE",
  "message": "ERROR_MESSAGE"
}
```

## Error Code

| Status code |          Code           |         Description          |
| :---------: | :---------------------: | :--------------------------- |
|     400     |      `BAD_REQUEST`      | リクエストが不正である       |
|     401     |     `UNAUTHORIZED`      | 認証が必要である             |
|     404     |       `NOT_FOUND`       | リソースが存在しない         |
|     409     |       `CONFLICT`        | リソースの状態が競合している |
|     500     | `INTERNAL_SERVER_ERROR` | サーバ内部エラー             |
|     503     |  `SERVICE_UNAVAILABLE`  | サービスが利用不可である     |
