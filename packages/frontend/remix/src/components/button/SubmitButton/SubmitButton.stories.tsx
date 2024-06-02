import type { Story, StoryDefault } from "@ladle/react";

import { SubmitButton } from ".";

export default {
  title: "button/SubmitButton",
} satisfies StoryDefault;

export const sample1: Story = () => {
  return <SubmitButton>Submit Button</SubmitButton>;
};
