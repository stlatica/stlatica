import { UserIcon } from "./UserIcon";
import { UserIconSampleImage } from "./UserIconSampleImage";

import type { Story, StoryDefault } from "@ladle/react";

export default {
  title: "common/UserIcon",
} satisfies StoryDefault;

export const Story1: Story = () => {
  return (
    <div style={{ width: "100px" }}>
      <UserIcon iconImage={UserIconSampleImage} />
    </div>
  );
};
Story1.storyName = "display icon 100px";

export const Story2: Story = () => {
  return (
    <div style={{ width: "300px" }}>
      <UserIcon iconImage={UserIconSampleImage} />
    </div>
  );
};
Story2.storyName = "display icon 300px";

export const Story3: Story = () => {
  return (
    <div style={{ width: "100%" }}>
      <UserIcon iconImage={UserIconSampleImage} />
    </div>
  );
};
Story3.storyName = "display icon 全幅";
