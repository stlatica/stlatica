// useNavigate が ladle 等からの利用時にエラーを起こしてしまうため、これで回避する
// https://github.com/remix-run/remix/discussions/9190
// 上記ディスカッションの内容を元に迂回策が提示されれば削除してよい

import { useNavigate } from "@remix-run/react";

/**
 * A version of `useNavigate` that works fine inside stories and other contexts
 * where Remix's router isn't available. Calling the original `useNavigate` in
 * those contexts throws an error.
 *
 * ⚠️ __Important__ If you use this function outside of Remix, such as in a
 * story, then the navigation will not actually happen.
 */
export const useNavigateSafe = () => {
  try {
    return useNavigate();
  } catch {
    return () => {};
  }
};
