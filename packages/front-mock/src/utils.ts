/**
 * 0~max-1 までの長さで配列を生成する
 * @param max
 * @returns 配列の要素番号と同じ数値が格納された配列
 */
export const InitArray = (max: number): number[] => {
  return new Array(Math.floor(Math.random() * max)).fill(0).map((x, i) => i);
};
