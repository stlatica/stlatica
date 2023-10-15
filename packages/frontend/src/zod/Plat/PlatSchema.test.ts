import * as Z from "./PlatSchema";
import { describe, expect, test } from "vitest";

describe("通常", () => {
  test("OK", () => {
    const result = Z.PlatSchema.safeParse({ plat: "test" });
    expect(result.success).toBe(true);
  });

  test("字数オーバー", () => {
    const result = Z.PlatSchema.safeParse({
      plat: "文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数",
    });
    expect(result.success).toBe(false);
  });
});

describe("境界値", () => {
  test("140字", () => {
    const result = Z.PlatSchema.safeParse({
      plat: "あいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこ",
    });
    expect(result.success).toBe(true);
  });

  test("141字", () => {
    const result = Z.PlatSchema.safeParse({
      plat: "あいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこ超",
    });
    expect(result.success).toBe(false);
  });
});
