import React from "react";

import { TextEditor } from "./TextEditor";

import type { Story, StoryDefault } from "@ladle/react";

export default {
  title: "components/common/TextEditor",
} satisfies StoryDefault;

export const Story1: Story = () => {
  return (
    <div className="h-32 bg-gray-800 p-6 text-white">
      <TextEditor componentId="story01" defaultValue="" maxlength={16} title="Name" />
    </div>
  );
};
Story1.storyName = "Name-[none]";

export const Story2: Story = () => {
  return (
    <div className="h-32 bg-gray-800 p-6 text-white">
      <TextEditor componentId="story02" defaultValue="nobuchi" maxlength={16} title="Name" />
    </div>
  );
};
Story2.storyName = "Name-[nobuchi]";

export const Story3: Story = () => {
  return (
    <div className="h-96 bg-gray-800 p-6 text-white">
      <TextEditor componentId="story03" defaultValue="" maxlength={512} title="Self-introduction" />
    </div>
  );
};
Story3.storyName = "self-introduction-[none]";

export const Story4: Story = () => {
  return (
    <div className="h-96 bg-gray-800 p-6 text-white">
      <TextEditor
        componentId="story04"
        defaultValue="CTRL OB(13期) 会計でした"
        maxlength={512}
        title="Self-introduction"
      />
    </div>
  );
};
Story4.storyName = "self-introduction-[nobuchi]";

export const Story5: Story = () => {
  return (
    <div className="h-96 bg-gray-200 p-6 text-black">
      <TextEditor
        componentId="story05"
        defaultValue="親要素の背景色と文字色はcomponentのtextarea, label要素に引き継がれない"
        maxlength={512}
        title="Self-introduction"
      />
    </div>
  );
};
Story5.storyName = "背景色は引き継げない";
