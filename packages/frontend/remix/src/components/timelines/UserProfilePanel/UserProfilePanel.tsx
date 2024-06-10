import { Title, Text, Paper, ActionIcon } from "@mantine/core";
import React from "react";
import { TbDots } from "react-icons/tb";

import { FollowButton } from "@/components/button/FollowButton";
import { UserIcon } from "@/components/common/UserIcon";

import {
  FollowsLine,
  container,
  header,
  icon,
  parent,
  texts,
  toolPalette,
} from "./UserProfilePanel.css";

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
      <img className={header} alt="header" src="https://via.placeholder.com/700/79b74a/fff.webp" />

      <div className={parent}>
        <div className={icon}>
          <UserIcon iconImage="https://via.placeholder.com/100/b7794a/fff.webp" />
        </div>
      </div>

      <div className={toolPalette}>
        <ActionIcon variant="outline" color="gray" radius="xl">
          <TbDots />
        </ActionIcon>
        <FollowButton isFollow={false} />
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
