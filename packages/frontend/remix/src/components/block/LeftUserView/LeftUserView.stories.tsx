import { LeftUserView } from "./LeftUserView";

import type { Story, StoryDefault } from "@ladle/react";

export default {
  title: "block/LeftUserView",
} satisfies StoryDefault;

export const Story1: Story = () => {
  return (
    <div className="h-[70vh] w-[300px]">
      <LeftUserView userID="userid" />
    </div>
  );
};
