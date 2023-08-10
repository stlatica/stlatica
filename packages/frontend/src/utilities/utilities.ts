/* eslint-disable no-console */

// TODO: ちょっと大きすぎになってきたので、ファイル分割等を検討中。

/**
 * Nullを判定します
 * @param x - Null型を除外したい値
 * @returns null | undefined だったら true
 * @example
 * const hoge: null | undefined | number
 *
 * if(IsNull(hoge))
 *     const x = hoge; // null | undefined
 * else
 *     const y = hoge; // numberであることが確定する
 */
export const IsNull = (x: unknown): x is null | undefined => {
  // return !IsNonNull(x);
  if (x === null || x === undefined) {
    return true;
  }
  return false;
};

/**
 * 配列の内容がユニークかチェックします
 * @param value - 配列
 * @returns 項目が重複していたらtrue
 */
export const IsArrayRepeated = (value: unknown[]) => {
  const x = new Set(value);

  return value.length !== x.size;
};

/**
 * 重い処理の再現用
 * @remarks デバッグ用なので、最終的にはコメントアウトで
 * @param delay - 待機時間 ミリ秒
 */
export const Sleep = (delay = 1000) => {
  return new Promise<void>((resolve) => {
    setTimeout(() => {
      resolve();
    }, delay);
  });
};

/**
 * 小数を丸めます
 * @param x - 数値
 * @param d - 桁数
 * @returns 数値
 * @example
 * // 小数第二位を四捨五入した文字列の取得例
 * MathRound(1.234, 1).toFix(1)
 */
export const MathRound = (x: number, d = 0) => {
  const digit = 10 ** d;

  return Math.round(x * digit) / digit;
};

/**
 * 小数を切り捨てします
 * @param x - 数値
 * @param d - 有効桁数
 * @returns 数値
 * @example
 * // 小数第二位を切り捨てした文字列の取得例
 * MathFloor(1.234, 1).toFix(1)
 */
export const MathFloor = (x: number, d = 0) => {
  const digit = 10 ** d;

  return Math.floor(x * digit) / digit;
};

/**
 * 小数を切り上げします
 * @param x - 数値
 * @param d - 有効桁数
 * @returns 数値
 * @example
 * // 小数第二位を切り上げした文字列の取得例
 * MathFloor(1.234, 1).toFix(1)
 */
export const MathCeil = (x: number, d = 0) => {
  const digit = 10 ** d;

  return Math.ceil(x * digit) / digit;
};

/**
 * 配列を前後に分けます
 * @param array -
 * @param maxLength -
 * @returns 2つの配列のタプル
 */
export const ArrayDevide = <T>(array: T[], maxLength: number) => {
  const result: [T[], T[]] = [array.slice(0, maxLength), array.slice(maxLength)];

  return result;
};

/**
 * 配列からのnull除外用
 * @param value
 * @returns
 */
export const nonNullableFilter = <T>(value: T): value is NonNullable<T> => {
  return value != null;
};

/**
 * 配列を一定サイズまでの長さに統一し、別々の配列に分解します
 * @example
 * [1,2,3,4,5,6,7,8,9,10], length=3
 * - [[1,2,3], [4,5,6], [7,8,9], [10]]
 * @example
 * [1,2,3,4,5], length=99
 * - [[1,2,3,4,5]]
 * @param array - 配列
 * @param maxLength - 1配列あたりの長さ
 * @returns 配列の配列
 */
export const ArrayFlex: <T>(array: T[], maxLength: number) => T[][] = <T>(
  array: T[],
  maxLength: number
) => {
  if (maxLength === 0) {
    console.warn("error ArrayFlex: maxLength=0. スライスせずに返します");
    return [array];
  }

  if (array.length <= maxLength) {
    return [array];
  }

  const [left, right] = ArrayDevide(array, maxLength);

  return [left, ...ArrayFlex(right, maxLength)];
};

/**
 * 特定の値を繰り返して初期化した配列を作成します
 * @param dataArray -
 * @param length -
 * @returns dataArrayをlengthの長さ分繰り返した配列
 * @throws dataArrayが空だった場合は例外を射出します
 */
export const ArrayCreateRepeat: <T>(dataArray: readonly T[], length: number) => T[] = (
  dataArray,
  length
) => {
  if (dataArray.length === 0) {
    throw new Error("ArrayCreateRepeat: dataArray length is 0");
  }

  const result = new Array(length)
    .fill(0)
    .map((_, i) => {
      return dataArray[i % dataArray.length];
    })
    .filter(nonNullableFilter);

  return result;
};

/**
 * チルドレンの型付け用
 */
export type OnlyChildren = {
  children?: React.ReactNode;
};

/**
 * 文字列中から空白を取り除きます
 * @param str
 * @returns
 */
export const RemoveWhiteSpace = (str: string) => {
  return str.replace(/\s/g, "");
};

/**
 * オブジェクトのUndefinedプロパティのみ除外する
 */
export const RemoveUndefinedAttributes: <T>(
  x: Record<string, T>
) => Record<string, NonNullable<T>> = (object) => {
  return Object.entries(object).reduce((prev, x) => {
    const key = x[0];
    const value = x[1];

    // nullだったら除外
    if (value === null || value === undefined) {
      return prev;
    }

    const current = { [key]: value };

    // 存在してたら追加して返す
    const ret = { ...prev, ...current };
    return ret;
  }, {});
};

/**
 * プリミティブ型配列のソート用
 * @param key オブジェクトのキー
 * @param reverse 反転するならtrue
 * @return -1, 1, 0
 * @example
 * [1,5,2,4].
 */
export const SimpleSort = <T>(a: T, b: T, reverse?: undefined | "reverse") => {
  const coe = reverse === "reverse" ? -1 : 1;
  if (a < b) {
    return -1 * coe;
  }
  if (a > b) {
    return 1 * coe;
  }
  return 0;
};

type KeysType<T extends Record<string, unknown>> = keyof T;

/**
 * オブジェクトを便利にソートします
 * @param key オブジェクトのキー
 * @param reverse 反転するならtrue
 * @return フィルタ用関数
 * @example
 * [{ x: 1 }, { x: 5 }, { x: 2 }].sort(SimpleSortObject("x"));
 */
export const SimpleSortObject = <T extends Record<string, unknown>>(
  key: KeysType<T>,
  reverse?: undefined | "reverse"
) => {
  return (a: T, b: T) => {
    return SimpleSort(a[key], b[key], reverse);
  };
};

/**
 * min, maxの範囲内に値を納めます
 */
export const Clamp = (value: number, min: number, max: number) => {
  return Math.max(min, Math.min(max, value));
};

/**
 * aタグを生成し、踏ませたことにしてダウンロードさせます
 * オリジンが違う場合は適用不能です
 * @param url URL
 * @param filename ファイル名を指定
 */
export const DownloadFromURL = (url: string, filename = "text.txt") => {
  const link = document.createElement("a");
  link.href = url;
  link.download = filename;
  document.body.appendChild(link);
  link.click();
  document.body.removeChild(link);
};
