import React from "react";

import { UserIconBase64Sample1 } from "./sample/UserIconBase64Sample1";
import { UserIcon } from "./UserIcon";

import type { Story, StoryDefault } from "@ladle/react";

export default {
  title: "components/common/UserIcon",
} satisfies StoryDefault;

export const Story1: Story = () => {
  return <UserIcon iconImage={UserIconBase64Sample1} size={100} />;
};
Story1.storyName = "size 100";

export const Story2: Story = () => {
  return <UserIcon iconImage={UserIconBase64Sample1} size={100} />;
};
Story1.storyName = "size 300";
