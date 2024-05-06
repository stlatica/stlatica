import { ProfileTimeline } from "./Timeline";

import type { Story, StoryDefault } from "@ladle/react";

export default {
  title: "timelines/Timeline",
} satisfies StoryDefault;

export const Story2: Story = () => {
  return <ProfileTimeline url="story/timeline/story2" user_id="sample_user" />;
};
Story2.storyName = "Profile Timeline";
