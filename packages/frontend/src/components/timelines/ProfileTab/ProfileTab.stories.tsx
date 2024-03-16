import React from "react";

import { ProfileTab } from "./ProfileTab";

import type { Story, StoryDefault } from "@ladle/react";

export default {
  title: "components/others/ProfileTab",
} satisfies StoryDefault;

export const Story1: Story = () => {
  return <ProfileTab />;
};
