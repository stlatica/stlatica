import { style } from "@vanilla-extract/css";

export const flexContainer = style({
  display: "flex",
});

export const flexColumn = style({
  display: "flex",
  width: "10rem", // 40 * 0.25rem (1rem = 16px)
  flexDirection: "column",
});
