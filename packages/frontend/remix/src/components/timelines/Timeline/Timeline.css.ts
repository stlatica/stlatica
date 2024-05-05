import { style } from "@vanilla-extract/css";

import { vars } from "@/styles/mantine";

export const container = style({
  overflowY: "scroll",
  height: "100%",
  gap: vars.spacing.xs,
});
