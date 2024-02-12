import React, { useEffect } from "react";
import { ThemeProvider } from "@mui/material/styles";

import { useLadleContext, ActionType, ThemeState } from "@ladle/react";
import type { GlobalProvider } from "@ladle/react";
import "../src/app/global.css";
import { muiTheme } from "../src/app/providers";

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
      <ThemeProvider theme={muiTheme}>{children}</ThemeProvider>
    </>
  );
};
