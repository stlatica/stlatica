import React from "react";
import type { Story, StoryDefault } from "@ladle/react";
import { PlatCell } from "./PlatCell";

export default {
  title: "components/common/PlatCell",
} satisfies StoryDefault;

export const Story1: Story = () => {
  return (
    <PlatCell
      userName="夏ノブチ"
      userId="@Nobuchi32384"
      content="3週連続でレート下がってるんだが。"
      replyCount={0}
      shareCount={1}
      favoriteCount={4}
    />
  );
};
Story1.storyName = "1行のPlat";

export const Story2: Story = () => {
  return (
    <PlatCell
      userName="冬ノブチ"
      userId="@Nobuchi79323"
      content="nobuchiさんのパナソニックグループ プログラミングコンテスト2023（AtCoder Beginner Contest 326）での成績：4526位, パフォーマンス：670相当, レーティング：789→777 (-12) "
      replyCount={1}
      shareCount={3}
      favoriteCount={8}
    />
  );
};
Story2.storyName = "複数行のPlat";
