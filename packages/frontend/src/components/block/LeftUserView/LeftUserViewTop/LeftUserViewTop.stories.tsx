import React from "react";

import { LeftUserViewTop } from "./LeftUserViewTop";

import type { Story, StoryDefault } from "@ladle/react";

export default {
  title: "components/block/LeftUserViewTop",
} satisfies StoryDefault;

export const Story1: Story = () => {
  return <LeftUserViewTop />;
};
