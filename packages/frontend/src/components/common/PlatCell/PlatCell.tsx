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
    <div className={styles.container}>
      {/* ユーザプロファイル */}
      <div className={styles.userProfile}>
        {/* アイコン */}
        <div className={styles.userIcon}>
          <UserIcon iconImage={userIcon} />
        </div>
        {/* ユーザ名 と ユーザID */}
        <div className={styles.userDetails}>
          {/* ユーザ名 */}
          <div className={styles.userName}>{userName}</div>
          {/* ユーザID */}
          <div className={styles.userId}>{userId}</div>
        </div>
        {/* プルダウンメニュー */}
        <div className={styles.menuButton}>
          <PlatMenuButton />
        </div>
      </div>
      {/* platの内容 */}
      <div className={styles.contentStyle}>{content}</div>
      {/* 投稿日時 */}
      <div>
        <div>{formatter.format(new Date(date))}</div>
      </div>
      {/* 下部のアイコン */}
      <div className={styles.bottomIcons}>
        {/* 返信数 */}
        <div className={styles.iconContainer}>
          <div className={styles.iconPadding}>
            <MdComment size="2em" />
          </div>
          <div>{replyCount}</div>
        </div>
        {/* シェア数 */}
        <div className={styles.iconContainer}>
          <div className={styles.iconPadding}>
            <MdAutorenew size="2em" />
          </div>
          <div>{shareCount}</div>
        </div>
        {/* いいね数 */}
        <div className={styles.iconContainer}>
          <div className={styles.iconPadding}>
            <MdStarBorder size="2em" />
          </div>
          <div>{favoriteCount}</div>
        </div>
      </div>
    </div>
  );
};
