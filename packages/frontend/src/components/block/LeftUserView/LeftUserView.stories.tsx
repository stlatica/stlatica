import React from "react";

import { LeftUserView } from "./LeftUserView";

import type { Story, StoryDefault } from "@ladle/react";

export default {
  title: "components/others/LeftUserView",
} satisfies StoryDefault;

export const Story1: Story = () => {
  return (
    <div className="h-[70vh] w-[400px]">
      <LeftUserView userID="userid" />
    </div>
  );
};
