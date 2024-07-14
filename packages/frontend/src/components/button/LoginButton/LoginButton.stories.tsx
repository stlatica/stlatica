import { LoginButton } from "./LoginButton";

import type { Story, StoryDefault } from "@ladle/react";

export default {
  title: "button/LoginButton",
} satisfies StoryDefault;

export const Story1: Story = () => {
  return (
    <LoginButton
      onClick={() => {
        console.log("clicked login button");
      }}
    />
  );
};
Story1.storyName = "login button";
