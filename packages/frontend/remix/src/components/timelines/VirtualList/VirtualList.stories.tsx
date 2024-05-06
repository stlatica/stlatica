import { VirtualList } from "./VirtualList";

import type { Story, StoryDefault } from "@ladle/react";

export default {
  title: "timelines/VirtualList",
} satisfies StoryDefault;

export const Story2: Story = () => {
  return (
    <VirtualList
      elements={[
        <div style={{ height: 100, backgroundColor: "red" }}>a</div>,
        <div style={{ height: 50, backgroundColor: "red" }}>a</div>,
        <div style={{ height: 20, backgroundColor: "red" }}>a</div>,
      ]}
    />
  );
};
Story2.storyName = "Profile Timeline";
