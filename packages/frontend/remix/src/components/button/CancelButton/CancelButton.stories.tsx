import { CancelButton } from "./CancelButton";

import type { Story, StoryDefault } from "@ladle/react";

export default {
  title: "button/CancelButton",
} satisfies StoryDefault;

export const Sample1: Story = () => {
  return (
    <div>
      <CancelButton>Cancel Button</CancelButton>
    </div>
  );
};
