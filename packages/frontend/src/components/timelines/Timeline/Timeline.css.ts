import { style } from "@vanilla-extract/css";

import { vars } from "@/styles/mantine.css";

export const container = style({
  overflowY: "scroll",
  height: "100%",
  contain: "strict",
});

export const scrollArea = style({
  width: "100%",
  position: "relative",
});

export const positionHelper = style({
  position: "absolute",
  top: 0,
  left: 0,
  width: "100%",
});

export const item = style({
  paddingTop: vars.spacing.xs,
});
