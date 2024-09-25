import * as v from "valibot";
import { describe, expect, test } from "vitest";

import * as Z from "./PlatSchema";

describe("通常", () => {
  test("OK", () => {
    const result = v.safeParse(Z.PlatSchema, { plat: "test" });
    expect(result.success).toBe(true);
  });

  test("字数オーバー", () => {
    const result = v.safeParse(Z.PlatSchema, {
      plat: "文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数",
    });
    expect(result.success).toBe(false);
  });
});

describe("境界値", () => {
  test("140字", () => {
    const result = v.safeParse(Z.PlatSchema, {
      plat: "あいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこ",
    });
    expect(result.success).toBe(true);
  });

  test("141字", () => {
    const result = v.safeParse(Z.PlatSchema, {
      plat: "あいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこあいうえおかきくけこ超",
    });
    expect(result.success).toBe(false);
  });
});
