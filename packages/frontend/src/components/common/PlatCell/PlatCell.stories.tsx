import { UserIconSampleImage } from "@/components/common/UserIcon/UserIconSampleImage";

import { PlatCell } from "./PlatCell";

import type { Story, StoryDefault } from "@ladle/react";

export default {
  title: "common/PlatCell",
} satisfies StoryDefault;

export const Story1: Story = () => {
  return (
    <PlatCell
      content="3週連続でレート下がってるんだが。"
      date="2021-01-30T00:00:00+09:00"
      favoriteCount={4}
      replyCount={0}
      shareCount={1}
      userIcon={UserIconSampleImage}
      userId="@Nobuchi32384"
      userName="夏ノブチ"
    />
  );
};
Story1.storyName = "1行のPlat";

export const Story2: Story = () => {
  return (
    <PlatCell
      content="nobuchiさんのパナソニックグループ プログラミングコンテスト2023（AtCoder Beginner Contest 326）での成績：4526位, パフォーマンス：670相当, レーティング：789→777 (-12) "
      date="2021-01-30T00:00:00+09:00"
      favoriteCount={8}
      replyCount={1}
      shareCount={3}
      userIcon={UserIconSampleImage}
      userId="@Nobuchi79323"
      userName="冬ノブチ"
    />
  );
};
Story2.storyName = "複数行のPlat";

export const Story3: Story = () => {
  return (
    <div style={{ width: "310px" }}>
      <PlatCell
        content="nobuchiさんのパナソニックグループ プログラミングコンテスト2023（AtCoder Beginner Contest 326）での成績：4526位, パフォーマンス：670相当, レーティング：789→777 (-12) "
        date="2021-01-30T00:00:00+09:00"
        favoriteCount={8}
        replyCount={1}
        shareCount={3}
        userIcon={UserIconSampleImage}
        userId="@Nobuchi79323"
        userName="冬ノブチ"
      />
    </div>
  );
};
Story3.storyName = "width=310";

export const Story4: Story = () => {
  return (
    <div style={{ width: "700px" }}>
      <PlatCell
        content="nobuchiさんのパナソニックグループ プログラミングコンテスト2023（AtCoder Beginner Contest 326）での成績：4526位, パフォーマンス：670相当, レーティング：789→777 (-12) "
        date="2021-01-30T00:00:00+09:00"
        favoriteCount={8}
        replyCount={1}
        shareCount={3}
        userIcon={UserIconSampleImage}
        userId="@Nobuchi79323"
        userName="冬ノブチ"
      />
    </div>
  );
};
Story4.storyName = "width=700";

export const Story5: Story = () => {
  return (
    <div style={{ width: "1000px" }}>
      <PlatCell
        content="nobuchiさんのパナソニックグループ プログラミングコンテスト2023（AtCoder Beginner Contest 326）での成績：4526位, パフォーマンス：670相当, レーティング：789→777 (-12) "
        date="2021-01-30T00:00:00+09:00"
        favoriteCount={8}
        replyCount={1}
        shareCount={3}
        userIcon={UserIconSampleImage}
        userId="@Nobuchi79323"
        userName="冬ノブチ"
      />
    </div>
  );
};
Story5.storyName = "width=1000";
