import { TextEditor } from "./TextEditor";

import type { Story, StoryDefault } from "@ladle/react";

export default {
  title: "common/TextEditor",
} satisfies StoryDefault;

export const Story1: Story = () => {
  return (
    <div style={{ height: "8rem" }}>
      <TextEditor defaultValue="" label="Name" maxLength={16} />
    </div>
  );
};
Story1.storyName = "Name-[none]";

export const Story2: Story = () => {
  return (
    <div style={{ height: "8rem" }}>
      <TextEditor defaultValue="nobuchi" label="Name" maxLength={16} />
    </div>
  );
};
Story2.storyName = "Name-[with lavel]";

export const Story3: Story = () => {
  return (
    <div style={{ height: "24rem" }}>
      <TextEditor defaultValue="" label="Self-introduction" maxLength={512} />
    </div>
  );
};
Story3.storyName = "self-introduction-[none]";

export const Story4: Story = () => {
  return (
    <div style={{ height: "24rem" }}>
      <TextEditor
        defaultValue="CTRL OB(13期) 会計でした"
        label="Self-introduction"
        maxLength={512}
      />
    </div>
  );
};
Story4.storyName = "self-introduction-[with lavel]";
