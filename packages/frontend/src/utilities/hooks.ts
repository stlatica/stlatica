import { useEffect } from "react";

/**
 * componentDidmountフック
 *
 * @remarks useEffect だとやれることが多すぎるので、これを使う方が可読性が高くおすすめです。
 * @param func 処理 () => void
 */
export const useComponentDidMount = (func: () => void) => {
  return useEffect(() => {
    return func();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);
};
