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
    <div style={{ padding: "20px", backgroundColor: "#373737", color: "#FFFFFF" }}>
      {/* ユーザプロファイル */}
      <div style={{ display: "flex", height: "50px" }}>
        {/* アイコン */}
        <div
          style={{
            height: "50px",
            width: "50px",
            backgroundColor: "#FFFFFF",
            color: "#000000",
            borderRadius: "100%",
          }}
        />
        <div style={{ paddingLeft: "13px", paddingRight: "13px" }}>
          {/* ユーザ名 */}
          <div style={{ fontSize: "100%", padding: "2%" }}>{userName}</div>
          {/* ユーザID */}
          <div style={{ fontSize: "75%", padding: "2%", color: "#C5C5C5" }}>{userId}</div>
        </div>
      </div>

      {/* platの内容 */}
      <div style={{ padding: "20px" }}>{content}</div>

      {/* 下部のアイコン */}
      <div style={{ display: "flex" }}>
        {/* 返信数 */}
        <div style={{ display: "flex", alignItems: "center", width: "100%" }}>
          <div className="px-2">
            <MdComment size="2em" />
          </div>
          <div>{replyCount}</div>
        </div>
        {/* シェア数 */}
        <div style={{ display: "flex", alignItems: "center", width: "100%" }}>
          <div className="px-2">
            <MdAutorenew size="2em" />
          </div>
          <div>{shareCount}</div>
        </div>
        {/* いいね数 */}
        <div style={{ display: "flex", alignItems: "center", width: "100%" }}>
          <div className="px-2">
            <MdStarBorder size="2em" />
          </div>
          <div>{favoriteCount}</div>
        </div>
      </div>
    </div>
  );
};
