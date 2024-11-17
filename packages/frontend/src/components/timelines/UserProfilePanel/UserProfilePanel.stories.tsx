import { UserProfilePanel } from "./UserProfilePanel";

import type { Story, StoryDefault } from "@ladle/react";

export default {
  title: "timelines/UserProfilePanel",
} satisfies StoryDefault;

export const Story1: Story = () => {
  return (
    <div style={{ display: "flex", justifyContent: "center", minWidth: "300px" }}>
      <div style={{ minWidth: "300px", maxWidth: "700px" }}>
        <UserProfilePanel userID="userid" />
      </div>
    </div>
  );
};

export const Story2: Story = () => {
  return (
    <div style={{ display: "flex", justifyContent: "center", minWidth: "300px" }}>
      <div style={{ width: "300px" }}>
        <UserProfilePanel userID="userid" />
      </div>
    </div>
  );
};
Story2.storyName = "width 300px";

export const Story3: Story = () => {
  return (
    <div style={{ display: "flex", justifyContent: "center", minWidth: "700px" }}>
      <div style={{ width: "700px" }}>
        <UserProfilePanel userID="userid" />
      </div>
    </div>
  );
};
Story3.storyName = "width 700px";
