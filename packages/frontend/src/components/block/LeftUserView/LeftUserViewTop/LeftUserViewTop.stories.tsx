import { LeftUserViewTop } from "./LeftUserViewTop";

import type { Story, StoryDefault } from "@ladle/react";

export default {
  title: "block/LeftUserViewTop",
} satisfies StoryDefault;

export const Story1: Story = () => {
  return (
    <div class="w-[300px]">
      <LeftUserViewTop />
    </div>
  );
};
