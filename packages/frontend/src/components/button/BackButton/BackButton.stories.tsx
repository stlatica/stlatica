import React from "react";

import { BackButton } from "./BackButton";

import type { Story, StoryDefault } from "@ladle/react";

export default {
  title: "components/button/BackButton",
} satisfies StoryDefault;

export const Story1: Story = () => {
  return <BackButton />;
};
