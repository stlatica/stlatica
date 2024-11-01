import type React from "react";
import { MdAutorenew, MdComment, MdStarBorder } from "react-icons/md";

import { PlatMenuButton } from "@/components/button/PlatMenuButton";
import { UserIcon } from "@/components/common/UserIcon";

import * as styles from "./PlatCell.css";

const formatter = new Intl.DateTimeFormat("ja-JP", {
  dateStyle: "medium",
  timeStyle: "medium",
  timeZone: "Asia/Tokyo",
});

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
export const PlatCell: React.FC<PlatCellProps> = ({
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
    <div class={styles.container}>
      {/* ユーザプロファイル */}
      <div class={styles.userProfile}>
        {/* アイコン */}
        <div class={styles.userIcon}>
          <UserIcon iconImage={userIcon} />
        </div>
        {/* ユーザ名 と ユーザID */}
        <div class={styles.userDetails}>
          {/* ユーザ名 */}
          <div class={styles.userName}>{userName}</div>
          {/* ユーザID */}
          <div class={styles.userId}>{userId}</div>
        </div>
        {/* プルダウンメニュー */}
        <div class={styles.menuButton}>
          <PlatMenuButton />
        </div>
      </div>
      {/* platの内容 */}
      <div class={styles.contentStyle}>{content}</div>
      {/* 投稿日時 */}
      <div>
        <div>{formatter.format(new Date(date))}</div>
      </div>
      {/* 下部のアイコン */}
      <div class={styles.bottomIcons}>
        {/* 返信数 */}
        <div class={styles.iconContainer}>
          <div class={styles.iconPadding}>
            <MdComment size="2em" />
          </div>
          <div>{replyCount}</div>
        </div>
        {/* シェア数 */}
        <div class={styles.iconContainer}>
          <div class={styles.iconPadding}>
            <MdAutorenew size="2em" />
          </div>
          <div>{shareCount}</div>
        </div>
        {/* いいね数 */}
        <div class={styles.iconContainer}>
          <div class={styles.iconPadding}>
            <MdStarBorder size="2em" />
          </div>
          <div>{favoriteCount}</div>
        </div>
      </div>
    </div>
  );
};
