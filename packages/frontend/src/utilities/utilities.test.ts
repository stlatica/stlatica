import { describe, expect, test } from "vitest";

import * as Utils from "./utilities";

describe("数学関数", () => {
  test("MathCeil", () => {
    expect(Utils.MathCeil(12.3456)).toBe(13);
    expect(Utils.MathCeil(12.3456, 0)).toBe(13);
    expect(Utils.MathCeil(12.3456, 1)).toBe(12.4);
    expect(Utils.MathCeil(12.3456, -1)).toBe(20);
  });

  test("MathFloor", () => {
    expect(Utils.MathFloor(12.3456)).toBe(12);
    expect(Utils.MathFloor(12.3456, 0)).toBe(12);
    expect(Utils.MathFloor(12.3456, 1)).toBe(12.3);
    expect(Utils.MathFloor(12.3456, -1)).toBe(10);
  });

  test("MathRound", () => {
    expect(Utils.MathRound(12.3456)).toBe(12);
    expect(Utils.MathRound(12.3456, 0)).toBe(12);
    expect(Utils.MathRound(12.3456, 1)).toBe(12.3);
    expect(Utils.MathRound(12.3444, 2)).toBe(12.34);
    expect(Utils.MathRound(12.3456, 2)).toBe(12.35);
    expect(Utils.MathRound(12.3456, -1)).toBe(10);
  });
});

describe("条件判定系", () => {
  test("IsArrayRepeated", () => {
    expect(Utils.IsArrayRepeated([1, 2, 3, 4])).toBe(false);
    expect(Utils.IsArrayRepeated([1, 2, 2])).toBe(true);
    expect(Utils.IsArrayRepeated([])).toBe(false);
  });

  test("nullcheck", () => {
    expect(Utils.IsNull("")).toBe(false);
    expect(Utils.IsNull(0)).toBe(false);
    expect(Utils.IsNull(undefined)).toBe(true);
    expect(Utils.IsNull(null)).toBe(true);
  });
});

describe("配列操作系", () => {
  test("ArrayDevide", () => {
    // 普通に分けるとき
    expect(JSON.stringify(Utils.ArrayDevide([1, 2, 3, 4, 5], 3))).toEqual(
      JSON.stringify([
        [1, 2, 3],
        [4, 5],
      ])
    );
    // 一致する時
    expect(JSON.stringify(Utils.ArrayDevide([1, 2, 3, 4, 5], 5))).toEqual(
      JSON.stringify([[1, 2, 3, 4, 5], []])
    );
    // 配列より長いとき
    expect(JSON.stringify(Utils.ArrayDevide([1, 2, 3, 4, 5], 7))).toEqual(
      JSON.stringify([[1, 2, 3, 4, 5], []])
    );
    // 0
    expect(JSON.stringify(Utils.ArrayDevide([1, 2, 3, 4, 5], 0))).toEqual(
      JSON.stringify([[], [1, 2, 3, 4, 5]])
    );
    // 空配列
    expect(JSON.stringify(Utils.ArrayDevide([1, 2, 3, 4, 5], 0))).toEqual(
      JSON.stringify([[], [1, 2, 3, 4, 5]])
    );
  });

  test("ArrayFlex", () => {
    // 普通に分けるとき ちょっと余る
    expect(JSON.stringify(Utils.ArrayFlex([1, 2, 3, 4, 5, 6, 7, 8], 3))).toEqual(
      JSON.stringify([
        [1, 2, 3],
        [4, 5, 6],
        [7, 8],
      ])
    );
    // ぴったり分割できるとき
    expect(JSON.stringify(Utils.ArrayFlex([1, 2, 3, 4, 5, 6, 7, 8], 2))).toEqual(
      JSON.stringify([
        [1, 2],
        [3, 4],
        [5, 6],
        [7, 8],
      ])
    );
    // 一致する時
    expect(JSON.stringify(Utils.ArrayFlex([1, 2, 3], 3))).toEqual(JSON.stringify([[1, 2, 3]]));
    // 配列より長いとき
    expect(JSON.stringify(Utils.ArrayFlex([1, 2], 999))).toEqual(JSON.stringify([[1, 2]]));
    // 0
    expect(JSON.stringify(Utils.ArrayFlex([1, 2, 3], 0))).toEqual(JSON.stringify([[1, 2, 3]]));
  });

  test("ArrayCreateRepeat", () => {
    // 普通に
    expect(JSON.stringify(Utils.ArrayCreateRepeat(["val1", "val2"], 5))).toEqual(
      JSON.stringify(["val1", "val2", "val1", "val2", "val1"])
    );
    expect(JSON.stringify(Utils.ArrayCreateRepeat(["val1"], 3))).toEqual(
      JSON.stringify(["val1", "val1", "val1"])
    );
    // 配列が0(例外発生)
    expect(() => {
      return Utils.ArrayCreateRepeat([], 2);
    }).toThrow("ArrayCreateRepeat: dataArray length is 0");
    // 0
    expect(JSON.stringify(Utils.ArrayCreateRepeat(["val1", "val2"], 0))).toEqual(
      JSON.stringify([])
    );
  });
});

test("RemoveWhiteSpace", () => {
  expect(Utils.RemoveWhiteSpace("a b c d")).toBe("abcd");

  expect(Utils.RemoveWhiteSpace(" a b cd ")).toBe("abcd");

  expect(Utils.RemoveWhiteSpace("\ta\tb\tc\td\t")).toBe("abcd");
  expect(Utils.RemoveWhiteSpace("　a　b　c　d　")).toBe("abcd");
});

test("RemoveUndefinedAttributes", () => {
  expect(Utils.RemoveUndefinedAttributes({ hoge: "hoge" })).toEqual({ hoge: "hoge" });
  expect(Utils.RemoveUndefinedAttributes({ hoge: undefined })).not.toEqual({ hoge: "hoge" });
  expect(Utils.RemoveUndefinedAttributes({ hoge: undefined })).toEqual({});
  expect(Utils.RemoveUndefinedAttributes({ hoge: undefined, huga: "huga" })).toEqual({
    huga: "huga",
  });
});

test("SimpleSort", () => {
  const s = Utils.SimpleSort;

  expect(JSON.stringify([3, 4, 2, 5, 1].sort(s))).toBe(JSON.stringify([1, 2, 3, 4, 5]));

  expect(JSON.stringify([1, 1, 2, 5, 1].sort(s))).toBe(JSON.stringify([1, 1, 1, 2, 5]));
});

describe("SimpleSortObject", () => {
  const usi = Utils.SimpleSortObject;
  const jst = JSON.stringify;

  type T = { x: number; y: string };

  const a: T = { x: 2, y: "aa" };
  const b: T = { x: 4, y: "bb" };
  const c: T = { x: 6, y: "2" };
  const d: T = { x: 8, y: "bb" };

  test("case1", () => {
    const x: T[] = [b, d, a, c];
    const y: T[] = [a, b, c, d];

    expect(jst(x.sort(usi("x")))).toBe(jst(y));

    expect(jst(x.sort(usi("x", "reverse")))).toBe(jst(y.reverse()));
  });

  test("Is stable sort?", () => {
    const x: T[] = [a, d, a, d, d, d];

    expect(jst(x.sort(usi("y")))).toBe(jst(x));
  });
});

test("Clamp", () => {
  expect(Utils.Clamp(2, 1, 5)).toBe(2);
  expect(Utils.Clamp(1.1, 1, 5)).toBe(1.1);
  expect(Utils.Clamp(0, 1, 5)).toBe(1);
  expect(Utils.Clamp(9, 1, 5)).toBe(5);

  expect(Utils.Clamp(2, -5, 5)).toBe(2);
  expect(Utils.Clamp(-3, -5, 5)).toBe(-3);
  expect(Utils.Clamp(-9, -5, 5)).toBe(-5);

  expect(Utils.Clamp(-9, -5, -3)).toBe(-5);
  expect(Utils.Clamp(0, -5, -3)).toBe(-3);
});
