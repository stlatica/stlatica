// ./edit_profile.css.ts
import { style } from "@vanilla-extract/css";

export const mainContainer = style({
  display: "flex",
  backgroundColor: "#1f2937", // bg-gray-800
  padding: "3rem", // p-12
  color: "white",
});

export const halfWidth = style({
  width: "50%",
});

export const marginLeft = style({
  marginLeft: "3rem", // ml-12
});

export const marginTop = style({
  marginTop: "3rem", // mt-12
});

export const flexContainer = style({
  display: "flex",
  justifyContent: "space-around",
});

export const iconEditorContainer = style({
  position: "relative",
  borderRadius: "0.5rem", // rounded-lg
  border: "1px solid", // border
  padding: "0.75rem", // p-3
  borderColor: "#4b5563", // dark:border-gray-600
});

export const iconEditorLabel = style({
  position: "absolute",
  transform: "translate(-0.875rem, -1.5rem) scale(0.75)", // -translate-x-3.5 -translate-y-6 scale-75
  backgroundColor: "#1f2937", // bg-gray-800
  padding: "0 0.5rem", // px-2
  color: "#9ca3af", // text-gray-400
});

export const iconEditorContent = style({
  width: "100px",
  height: "100px",
  backgroundColor: "white",
  color: "black",
});
