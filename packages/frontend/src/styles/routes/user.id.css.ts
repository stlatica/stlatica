import { style } from "@vanilla-extract/css";

export const mainContainer = style({
  height: "100vh", // h-screen
});

export const flexContainer = style({
  display: "flex",
  height: "100%", // h-full
});

export const leftPanel = style({
  width: "330px",
  color: "#1e293b", // text-slate-900
});

export const rightPadding = style({
  paddingRight: "0.5rem", // pr-2
});
