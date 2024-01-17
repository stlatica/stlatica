import React from "react";

import { LoginScene } from "./LoginScene";

import type { Story, StoryDefault } from "@ladle/react";

export default {
  title: "components/others/LoginScene",
} satisfies StoryDefault;

export const Story1: Story = () => {
  return <LoginScene />;
};
