import { UserProfilePanel } from "./UserProfilePanel";

import type { Story, StoryDefault } from "@ladle/react";

export default {
  title: "timelines/User-Profile-Header",
} satisfies StoryDefault;

export const Story1: Story = () => {
  return (
    <div style={{ display: "flex", justifyContent: "center", width: "100%" }}>
      <div style={{ minWidth: "300px", maxWidth: "600px" }}>
        <UserProfilePanel userID="userid" />
      </div>
    </div>
  );
};
export const Story2: Story = () => {
  return (
    <div style={{ display: "flex", justifyContent: "center", width: "300px" }}>
      <div style={{ minWidth: "300px", maxWidth: "600px" }}>
        <UserProfilePanel userID="userid" />
      </div>
    </div>
  );
};
export const Story3: Story = () => {
  return (
    <div style={{ display: "flex", justifyContent: "center", width: "700px" }}>
      <div style={{ minWidth: "300px", maxWidth: "600px" }}>
        <UserProfilePanel userID="userid" />
      </div>
    </div>
  );
};
