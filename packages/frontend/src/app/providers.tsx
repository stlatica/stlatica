"use client";

import { createTheme } from "@mui/material/styles";

import { OnlyChildren } from "@/utilities/utilities";

export const muiTheme = createTheme({
  palette: {
    mode: "dark",
  },
  typography: {
    button: {
      textTransform: "none",
    },
  },
});

export const Providers: React.FC<OnlyChildren> = ({ children }) => {
  return <>{children}</>;
};
