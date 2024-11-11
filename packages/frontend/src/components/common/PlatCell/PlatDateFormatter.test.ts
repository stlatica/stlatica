import { describe, expect, test } from "vitest";

import { PlatDateFormatter } from "./PlatDateFormatter";

describe("相対時間の計算 PlatDateFormatter", () => {
  test("タイムスタンプが現在時刻を超過", () => {
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2020-01-01T00:00:00+09:00"),
        currentTime: new Date("2019-12-12T23:59:59+09:00"),
      }),
    ).toBe("現在");
  });

  test("タイムスタンプが現在時刻と一致", () => {
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2020-01-01T00:00:00+09:00"),
        currentTime: new Date("2020-01-01T00:00:00+09:00"),
      }),
    ).toBe("現在");
  });

  test("秒差", () => {
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2020-01-01T00:00:00+09:00"),
        currentTime: new Date("2020-01-01T00:00:01+09:00"),
      }),
    ).toBe("1 秒前");
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2020-01-01T00:00:00+09:00"),
        currentTime: new Date("2020-01-01T00:00:02+09:00"),
      }),
    ).toBe("2 秒前");
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2020-01-01T00:00:00+09:00"),
        currentTime: new Date("2020-01-01T00:00:59+09:00"),
      }),
    ).toBe("59 秒前");
  });

  test("分差", () => {
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2020-01-01T00:00:00+09:00"),
        currentTime: new Date("2020-01-01T00:01:00+09:00"),
      }),
    ).toBe("1 分前");
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2020-01-01T00:00:00+09:00"),
        currentTime: new Date("2020-01-01T00:02:00+09:00"),
      }),
    ).toBe("2 分前");
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2020-01-01T00:00:00+09:00"),
        currentTime: new Date("2020-01-01T00:59:00+09:00"),
      }),
    ).toBe("59 分前");
  });

  test("時間差", () => {
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2020-01-01T00:00:00+09:00"),
        currentTime: new Date("2020-01-01T01:00:00+09:00"),
      }),
    ).toBe("1 時間前");
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2020-01-01T00:00:00+09:00"),
        currentTime: new Date("2020-01-01T02:00:00+09:00"),
      }),
    ).toBe("2 時間前");
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2020-01-01T00:00:00+09:00"),
        currentTime: new Date("2020-01-01T23:00:00+09:00"),
      }),
    ).toBe("23 時間前");
  });

  test("日差", () => {
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2020-01-01T00:00:00+09:00"),
        currentTime: new Date("2020-01-02T00:00:00+09:00"),
      }),
    ).toBe("1 日前");
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2020-01-01T00:00:00+09:00"),
        currentTime: new Date("2020-01-03T00:00:00+09:00"),
      }),
    ).toBe("2 日前");
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2020-01-01T00:00:00+09:00"),
        currentTime: new Date("2020-01-30T00:00:00+09:00"),
      }),
    ).toBe("29 日前");
  });

  test("月差", () => {
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2020-01-01T00:00:00+09:00"),
        currentTime: new Date("2020-02-01T00:00:00+09:00"),
      }),
    ).toBe("1 ヶ月前");
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2020-01-01T00:00:00+09:00"),
        currentTime: new Date("2020-03-01T00:00:00+09:00"),
      }),
    ).toBe("2 ヶ月前");
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2020-01-01T00:00:00+09:00"),
        currentTime: new Date("2020-12-01T00:00:00+09:00"),
      }),
    ).toBe("11 ヶ月前");
  });

  test("年差", () => {
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2020-01-01T00:00:00+09:00"),
        currentTime: new Date("2021-01-01T00:00:00+09:00"),
      }),
    ).toBe("1 年前");
  });
});
