// @/styles/routes/login.css.ts
import { style } from "@vanilla-extract/css";

export const mainContainer = style({
  display: "flex",
  height: "100vh",
  width: "100vw",
  justifyContent: "center",
  backgroundColor: "#1f2937", // bg-gray-800
});

export const innerContainer = style({
  display: "flex",
  width: "80%",
  minWidth: "300px",
  maxWidth: "800px",
  flexDirection: "column",
  justifyContent: "center",
  gap: "1.5rem", // gap-6
});

export const title = style({
  display: "flex",
  justifyContent: "center",
  fontSize: "3rem", // text-5xl
});

export const textEditorContainer = style({
  backgroundColor: "#1f2937", // bg-gray-800
  color: "white",
});

export const buttonContainer = style({
  display: "flex",
  justifyContent: "center",
});
