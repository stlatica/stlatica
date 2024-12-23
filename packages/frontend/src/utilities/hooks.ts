import { useEffect } from "react";

/**
 * componentDidmountフック
 *
 * @remarks useEffect だとやれることが多すぎるので、マウント時のみの処理はこれを推奨します。
 * @param func unmountは処理しません
 */
export const useComponentDidMount = (func: () => void) => {
  // biome-ignore lint/correctness/useExhaustiveDependencies: <explanation>
  useEffect(() => {
    func();
  }, []);
};
