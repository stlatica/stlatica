import { ProfileTab } from "./ProfileTab";

import type { Story, StoryDefault } from "@ladle/react";

export default {
  title: "timelines/ProfileTab",
} satisfies StoryDefault;

export const Story1: Story = () => {
  return <ProfileTab />;
};
