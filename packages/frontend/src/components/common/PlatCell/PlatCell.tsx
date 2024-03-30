import { Menu } from "@mantine/core";
import React from "react";
import { MdAutorenew, MdComment, MdStarBorder, MdMoreHoriz } from "react-icons/md";

import { UserIcon } from "@/components/common/UserIcon";
import { DateTime } from "@/components/others/DateTime";

type PlatCellProps = {
  readonly content: string;
  readonly createdAt: string;
  readonly favoriteCount: number;
  readonly replyCount: number;
  readonly shareCount: number;
  readonly userIcon: string;
  readonly userId: string;
  readonly userName: string;
};

/**
 *
 */
export const PlatCell: React.FC<PlatCellProps> = ({
  content,
  createdAt,
  favoriteCount,
  replyCount,
  shareCount,
  userIcon,
  userId,
  userName,
}) => {
  const formatDateLocale = (date: Date): string => {
    return date.toLocaleString();
  };
  return (
    <div className="border-2 border-gray-500 bg-gray-800 p-5 text-white">
      {/* ユーザプロファイル */}
      <div className="flex h-12">
        {/* アイコン */}
        <div className="size-12">
          <UserIcon iconImage={userIcon} />
        </div>
        {/* ユーザ名 と ユーザID */}
        <div className="px-3">
          {/* ユーザ名 */}
          <div className="text-lg">{userName}</div>
          {/* ユーザID */}
          <div className="text-sm text-gray-500">{userId}</div>
        </div>
        {/* プルダウンメニュー    参考 : https://mantine.dev/core/menu/ */}
        <div className="ml-auto">
          <Menu>
            <Menu.Target>
              <MdMoreHoriz />
            </Menu.Target>
            <Menu.Dropdown>
              <Menu.Item>Delete this plat</Menu.Item>
              <Menu.Item>Mute this user</Menu.Item>
            </Menu.Dropdown>
          </Menu>
        </div>
      </div>
      {/* platの内容 */}
      <div className="p-5">{content}</div>
      {/* 投稿日時 */}
      <div>
        <DateTime date={createdAt} formatDate={formatDateLocale} />
      </div>
      {/* 下部のアイコン */}
      <div className="flex">
        {/* 返信数 */}
        <div className="flex flex-1 items-center">
          <div className="px-2">
            <MdComment size="2em" />
          </div>
          <div>{replyCount}</div>
        </div>
        {/* シェア数 */}
        <div className="flex flex-1 items-center">
          <div className="px-2">
            <MdAutorenew size="2em" />
          </div>
          <div>{shareCount}</div>
        </div>
        {/* いいね数 */}
        <div className="flex flex-1 items-center">
          <div className="px-2">
            <MdStarBorder size="2em" />
          </div>
          <div>{favoriteCount}</div>
        </div>
      </div>
    </div>
  );
};
