import { Timeline } from "./Timeline";

import type { Story, StoryDefault } from "@ladle/react";

export default {
  title: "components/timelines/Timeline",
} satisfies StoryDefault;

export const Story1: Story = () => {
  return <Timeline />;
};
