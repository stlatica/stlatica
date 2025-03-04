import { ActionIcon, Paper, Text, Title } from "@mantine/core";
import type React from "react";
import { TbDots } from "react-icons/tb";

import { FollowButton } from "@/components/button/FollowButton";
import { UserIcon } from "@/components/common/UserIcon";
import { SampleImageBrown400 } from "@/features/sample/sample-image/SampleImageBrown400";
import { SampleImageGreen400 } from "@/features/sample/sample-image/SampleImageGreen400";

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
};

/**
 * ユーザープロフィール画面での、ユーザー情報表示部
 */
export const UserProfilePanel: React.FC<LeftUserViewProps> = ({ userID }) => {
  // ストアからユーザーデータを引っ張ってくる

  return (
    <Paper className={container}>
      <img className={header} alt="header" src={SampleImageGreen400} />

      <div className={parent}>
        <div className={icon}>
          <UserIcon iconImage={SampleImageBrown400} />
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
