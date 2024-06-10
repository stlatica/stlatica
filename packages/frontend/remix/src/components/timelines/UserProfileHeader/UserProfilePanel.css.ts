import { style } from "@vanilla-extract/css";

export const container = style({
  width: "100%",
  height: "100%",
  // backgroundColor: "var(--mantine-color-gray-9)",
});

export const header = style({
  height: "200px",
  width: "100%",
  objectFit: "cover",
  display: "flex",
  justifyContent: "center",
  backgroundColor: "black",
});

const border = "3px";
const width = 100;

export const parent = style({
  position: "relative",
  height: `${width / 2}px`,
  marginBottom: "0.5rem",
  marginLeft: `${width / 5}px`,
});

export const icon = style({
  width: `${width}px`,
  padding: border,
  margin: border,
  backgroundColor: "black",
  borderRadius: "50%",
  position: "absolute",
  top: `-${width / 2}px`,
});

export const texts = style({
  margin: "1rem",
  display: "flex",
  flexDirection: "column",
});

export const FollowsLine = style({
  display: "flex",
  flexDirection: "row",
  alignItems: "center",
  gap: "1rem",
  marginTop: "0.5rem",
  marginBottom: "0.5rem",
});
