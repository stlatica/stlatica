import { style } from "@vanilla-extract/css";

export const fullHeight = style({
  height: "100vh", // h-screen
});

export const flexContainer = style({
  display: "flex",
  height: "100%", // h-full
  flexDirection: "row", // flex-row
  gap: "1rem", // gap-4
});

export const sidebar = style({
  display: "flex",
  height: "100%", // h-full
  width: "8em", // w-[8em]
  flexDirection: "column", // flex-col
  gap: "0.25rem", // gap-1
});
