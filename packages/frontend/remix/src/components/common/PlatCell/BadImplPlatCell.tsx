import React from "react";
import { MdAutorenew, MdComment, MdStarBorder } from "react-icons/md";

import { PlatMenuButton } from "@/components/button/PlatMenuButton";
import { UserIcon } from "@/components/common/UserIcon";

const dateObj = (dateStr: string): Date => {
  return new Date(dateStr);
};

// const formatter = new Intl.DateTimeFormat("ja-JP", {
//   dateStyle: "medium",
//   timeStyle: "medium",
//   timeZone: "Asia/Tokyo",
// });

// 毎回のフォーマットでインスタンスを生成
const convertDateNotReuseFormatter = (d: Date) => {
  return new Intl.DateTimeFormat("ja-JP", { dateStyle: "medium", timeStyle: "medium" }).format(d);
};

type PlatCellProps = {
  readonly content: string;
  readonly date: string;
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
export const BadImplPlatCell: React.FC<PlatCellProps> = ({
  content,
  date,
  favoriteCount,
  replyCount,
  shareCount,
  userIcon,
  userId,
  userName,
}) => {
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
          <PlatMenuButton />
        </div>
      </div>
      {/* platの内容 */}
      <div className="p-5">{content}</div>
      {/* 投稿日時 */}
      <div>
        <div>{convertDateNotReuseFormatter(dateObj(date))}</div>
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
