import { describe, expect, test } from "vitest";

import { PlatDateFormatter } from "./PlatDateFormatter";

describe("相対時間の計算 PlatDateFormatter", () => {
  test("不正な引数：タイムスタンプが現在時刻を超過", () => {
    expect(() => {
      PlatDateFormatter({
        platTimestamp: new Date("2021-01-30T00:00:55+09:00"),
        currentTime: new Date("2021-01-30T00:00:54+09:00"),
      });
    }).toThrow("Invalid time difference: platTimestamp is later than currentTime.");
  });

  test("秒差", () => {
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2021-01-30T00:00:55+09:00"),
        currentTime: new Date("2021-01-30T00:00:55+09:00"),
      })
    ).toBe("0 秒前");
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2021-01-30T00:00:55+09:00"),
        currentTime: new Date("2021-01-30T00:00:56+09:00"),
      })
    ).toBe("1 秒前");
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2021-01-30T00:00:55+09:00"),
        currentTime: new Date("2021-01-30T00:01:54+09:00"),
      })
    ).toBe("59 秒前");
  });

  test("分差", () => {
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2021-01-30T00:55:55+09:00"),
        currentTime: new Date("2021-01-30T00:56:55+09:00"),
      })
    ).toBe("1 分前");
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2021-01-30T00:55:55+09:00"),
        currentTime: new Date("2021-01-30T01:54:55+09:00"),
      })
    ).toBe("59 分前");
  });

  test("時間差", () => {
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2021-01-05T20:55:55+09:00"),
        currentTime: new Date("2021-01-05T21:55:55+09:00"),
      })
    ).toBe("1 時間前");
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2021-01-05T20:55:55+09:00"),
        currentTime: new Date("2021-01-06T19:55:55+09:00"),
      })
    ).toBe("23 時間前");
  });

  test("日差", () => {
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2021-01-15T20:55:55+09:00"),
        currentTime: new Date("2021-01-16T20:55:55+09:00"),
      })
    ).toBe("1 日前");
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2021-01-15T20:55:55+09:00"),
        currentTime: new Date("2021-02-13T20:55:55+09:00"),
      })
    ).toBe("29 日前");
  });

  test("月差", () => {
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2020-12-15T20:55:55+09:00"),
        currentTime: new Date("2021-01-14T20:55:55+09:00"),
      })
    ).toBe("1 ヶ月前");
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2020-02-15T20:55:55+09:00"),
        currentTime: new Date("2021-01-14T20:55:55+09:00"),
      })
    ).toBe("11 ヶ月前");
  });

  test("年差", () => {
    expect(
      PlatDateFormatter({
        platTimestamp: new Date("2020-01-15T20:55:55+09:00"),
        currentTime: new Date("2021-01-14T20:55:55+09:00"),
      })
    ).toBe("1 年前");
  });
});
