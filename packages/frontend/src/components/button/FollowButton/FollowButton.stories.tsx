import React from "react";
import type { Story, StoryDefault } from "@ladle/react";
import { FollowButton } from "./FollowButton";
// import FollowButton from "./FollowButton";

export default {
  title: "components/button/FollowButton",
} satisfies StoryDefault;

export const Story1: Story = () => {
  return <FollowButton initialFollowState={false} />;
};
