import { style } from "@vanilla-extract/css";

export const mainContainer = style({
  // height: "100vh", // h-screen
  height: "100%",

  display: "flex",
  justifyContent: "center",
});

export const flexContainer = style({
  // height: "100%", // h-full
  maxWidth: "720px",
});
