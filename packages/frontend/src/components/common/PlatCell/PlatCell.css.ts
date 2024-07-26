// import { style } from "@vanilla-extract/css";

// export const container = style({
//   border: "2px solid #ccc",
//   backgroundColor: "rgb(31 41 55)",
//   padding: "1.25rem",
//   color: "#fff",
// });

// export const userProfile = style({
//   display: "flex",
//   height: "3rem",
// });

import { style } from "@vanilla-extract/css";

export const container = style({
  border: "2px solid #6b7280", // gray-500
  backgroundColor: "#1f2937", // gray-800
  padding: "1.25rem", // p-5
  color: "#ffffff", // text-white
});

export const userProfile = style({
  display: "flex",
  height: "3rem", // h-12
});

export const userIcon = style({
  width: "3rem", // size-12
  height: "3rem", // size-12
});

export const userDetails = style({
  paddingLeft: "0.75rem", // px-3
});

export const userName = style({
  fontSize: "1.125rem", // text-lg
});

export const userId = style({
  fontSize: "0.875rem", // text-sm
  color: "#6b7280", // text-gray-500
});

export const menuButton = style({
  marginLeft: "auto", // ml-auto
});

export const contentStyle = style({
  padding: "1.25rem", // p-5
});

export const bottomIcons = style({
  display: "flex",
});

export const iconContainer = style({
  display: "flex",
  flex: 1,
  alignItems: "center",
});

export const iconPadding = style({
  paddingLeft: "0.5rem", // px-2
  paddingRight: "0.5rem", // px-2
});
