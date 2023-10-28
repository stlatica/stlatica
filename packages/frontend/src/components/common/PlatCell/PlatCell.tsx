import React from "react";
import CommentIcon from "@mui/icons-material/Comment";
import AutorenewIcon from "@mui/icons-material/Autorenew";
import StarBorderIcon from "@mui/icons-material/StarBorder";
import { Padding } from "@mui/icons-material";

type PlatCellProps = {
  readonly userName: string;
  readonly userId: string;
  readonly content: string;
  readonly replyCount: number;
  readonly shareCount: number;
  readonly favoriteCount: number;
};

/**
 *
 */
export const PlatCell: React.FC<PlatCellProps> = ({
  userName,
  userId,
  content,
  replyCount,
  shareCount,
  favoriteCount,
}) => {
  return (
    <div style={{ width: "310px", padding: "2%", backgroundColor: "#373737", color: "#FFFFFF" }}>
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
        ></div>
        <div style={{ paddingLeft: "2%", paddingRight: "2%" }}>
          {/* ユーザ名 */}
          <div style={{ fontSize: "100%", padding: "2%" }}>{userName}</div>
          {/* ユーザID */}
          <div style={{ fontSize: "75%", padding: "2%", color: "#C5C5C5" }}>{userId}</div>
        </div>
      </div>

      {/* platの内容 */}
      <div style={{ padding: "5%" }}>{content}</div>

      {/* 下部のアイコン */}
      <div style={{ display: "flex" }}>
        {/* 返信数 */}
        <div style={{ display: "flex", alignItems: "center", width: "100%" }}>
          <CommentIcon style={{ paddingLeft: "5%", paddingRight: "5%" }}></CommentIcon>
          <div>{replyCount}</div>
        </div>
        {/* シェア数 */}
        <div style={{ display: "flex", alignItems: "center", width: "100%" }}>
          <AutorenewIcon style={{ paddingLeft: "5%", paddingRight: "5%" }}></AutorenewIcon>
          <div>{shareCount}</div>
        </div>
        {/* いいね数 */}
        <div style={{ display: "flex", alignItems: "center", width: "100%" }}>
          <StarBorderIcon style={{ paddingLeft: "5%", paddingRight: "5%" }}></StarBorderIcon>
          <div>{favoriteCount}</div>
        </div>
      </div>
    </div>
  );
};
