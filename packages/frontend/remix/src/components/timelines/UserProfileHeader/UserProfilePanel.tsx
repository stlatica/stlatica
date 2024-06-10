import { Title, Text, Paper } from "@mantine/core";
import React from "react";

import { UserIcon } from "@/components/common/UserIcon";

import { FollowsLine, container, header, icon, parent, texts } from "./UserProfilePanel.css";

type LeftUserViewProps = {
  readonly userID: string;
  // children: React.ReactNode;
};

/**
 * ユーザープロフィール画面での、ユーザー情報表示部
 */
export const UserProfilePanel: React.FC<LeftUserViewProps> = ({ userID }) => {
  // ストアからユーザーデータを引っ張ってくる

  return (
    <Paper className={container}>
      <img
        className={header}
        alt="header"
        src="https://generative-placeholders.glitch.me/image?width=400&height=300&style=triangles&gap=100"
      />

      <div className={parent}>
        <div className={icon}>
          <UserIcon iconImage="https://generative-placeholders.glitch.me/image?width=300&height=300&style=mondrian" />
        </div>
      </div>

      <Paper className={texts}>
        <Title>User Name Hereeeeeeeee</Title>
        <Text>@{userID}</Text>
        <Paper className={FollowsLine}>
          <Text>100 follower</Text>
          <Text>100 follow</Text>
          <Text>100 post</Text>
        </Paper>
        <Text>
          ここに自己紹介をかきまああああああああああああああああああああああああああああああああああああす文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数
        </Text>
      </Paper>
    </Paper>
  );
};
