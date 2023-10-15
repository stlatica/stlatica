import React, { useState } from "react";
// eslint-disable-next-line import/no-extraneous-dependencies
import { Story, StoryDefault } from "@ladle/react";
import { FollowButton } from "./FollowButton";

export default {
  title: "components/button/FollowButton",
} satisfies StoryDefault;

export const Story1: Story = () => {
  return (
    <FollowButton
      isFollow={false}
      onClick={() => {
        console.log("push Story1");
      }}
    />
  );
};
Story1.storyName = "未フォロー";

export const Story2: Story = () => {
  return (
    <FollowButton
      isFollow
      onClick={() => {
        console.log("push Story2");
      }}
    />
  );
};
Story2.storyName = "フォロー中";

export const Story3: Story = () => {
  // 外部で保持している、ボタンの状態
  // 本来はAPIから取得するが、動作確認用にローカルで持っておく
  const [isFollow, setIsFollow] = useState(false);

  return (
    <FollowButton
      isFollow={isFollow}
      onClick={() => {
        // 押すたびに反転する
        setIsFollow(!isFollow);
      }}
    />
  );
};
Story3.storyName = "_動作サンプル";
