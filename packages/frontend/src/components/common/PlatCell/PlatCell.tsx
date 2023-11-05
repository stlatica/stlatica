import React from "react";
import { MdAutorenew, MdComment, MdStarBorder } from "react-icons/md";

type PlatCellProps = {
  readonly content: string;
  readonly favoriteCount: number;
  readonly replyCount: number;
  readonly shareCount: number;
  readonly userId: string;
  readonly userName: string;
};

/**
 *
 */
export const PlatCell: React.FC<PlatCellProps> = ({
  content,
  favoriteCount,
  replyCount,
  shareCount,
  userId,
  userName,
}) => {
  return (
    <div className="p-5 bg-gray-800 text-white">
      {/* ユーザプロファイル */}
      <div className="flex h-12">
        {/* アイコン */}
        <div className="h-12 w-12 bg-white text-black rounded-full" />
        {/* ユーザ名 と ユーザID */}
        <div className="px-3">
          {/* ユーザ名 */}
          <div className="text-lg">{userName}</div> {/* ユーザID */}
          <div className="text-sm text-gray-500">{userId}</div>{" "}
        </div>
      </div>
      {/* platの内容 */}
      <div className="p-5">{content}</div>
      {/* 下部のアイコン */}
      <div className="flex">
        {/* 返信数 */}
        <div className="flex items-center flex-1">
          <div className="px-2">
            <MdComment size="2em" />
          </div>
          <div>{replyCount}</div>
        </div>
        {/* シェア数 */}
        <div className="flex items-center flex-1">
          <div className="px-2">
            <MdAutorenew size="2em" />
          </div>
          <div>{shareCount}</div>
        </div>
        {/* いいね数 */}
        <div className="flex items-center flex-1">
          <div className="px-2">
            <MdStarBorder size="2em" />
          </div>
          <div>{favoriteCount}</div>
        </div>
      </div>
    </div>
  );
};
