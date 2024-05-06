import "../src/routes/global.css";
import "@mantine/core/styles.css";

import React, { useEffect } from "react";

import { useLadleContext, ActionType, ThemeState } from "@ladle/react";
import type { GlobalProvider } from "@ladle/react";
import { MantineProvider, useMantineColorScheme } from "@mantine/core";

import { Providers } from "../src/Providers";

/**
 * useMantineColorScheme は MantineProvider 以下でないと呼べないのでこうなる
 */
const MantineWrap: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const { setColorScheme } = useMantineColorScheme();

  const { globalState } = useLadleContext();

  useEffect(() => {
    setColorScheme(globalState.theme);
  }, [globalState.theme]);
  return <>{children}</>;
};

export const Provider: GlobalProvider = ({ children, globalState }) => {
  const { dispatch } = useLadleContext();

  useEffect(() => {
    // 強制ダークモード化
    dispatch({
      type: ActionType.UpdateTheme,
      value: ThemeState.Dark,
    });
  }, []);

  return (
    <>
      <Providers>
        <MantineWrap>{children}</MantineWrap>
      </Providers>
    </>
  );
};
