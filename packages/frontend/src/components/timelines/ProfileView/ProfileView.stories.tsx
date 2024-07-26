import { LeftUserView } from "./ProfileView";

import type { Story, StoryDefault } from "@ladle/react";

export default {
  title: "timelines/ProfileView",
} satisfies StoryDefault;

export const Story1: Story = () => {
  return (
    <div style={{ height: "70vh", width: "300px" }}>
      <LeftUserView userID="userid" />
    </div>
  );
};
