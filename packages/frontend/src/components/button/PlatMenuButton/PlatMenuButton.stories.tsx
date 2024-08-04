import type { Story, StoryDefault } from "@ladle/react";

import { PlatMenuButton } from ".";

export default {
  title: "button/PlatMenuButton",
} satisfies StoryDefault;

export const sample1: Story = () => {
  return <PlatMenuButton />;
};
