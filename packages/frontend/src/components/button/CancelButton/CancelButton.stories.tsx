import { Sample1 } from "./sample";

import type { Story, StoryDefault } from "@ladle/react";

export default {
  title: "components/button/CancelButton",
} satisfies StoryDefault;

export const sample1: Story = () => {
  return <Sample1 />;
};
