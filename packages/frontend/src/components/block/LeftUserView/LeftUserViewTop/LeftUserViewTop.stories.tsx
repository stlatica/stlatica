import React from "react";
import type { Story, StoryDefault } from "@ladle/react";
import { LeftUserViewTop } from "./LeftUserViewTop";

export default {
  title: "components/block/LeftUserViewTop",
} satisfies StoryDefault;

export const Story1: Story = () => {
  return <LeftUserViewTop />;
};
