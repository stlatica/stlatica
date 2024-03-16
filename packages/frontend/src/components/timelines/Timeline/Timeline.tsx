import React from "react";

import { PlatCell } from "@/components/common/PlatCell";

type TimelineProps = {
  // children: React.ReactNode;
};

/**
 *
 */
export const Timeline: React.FC<TimelineProps> = () => {
  return (
    <div className="flex flex-col gap-2">
      <PlatCell
        content="はじめてのplat"
        favoriteCount={0}
        replyCount={0}
        shareCount={0}
        userId="test"
        userName="テストユーザー"
      />
      <PlatCell
        content="はじめてのplat"
        favoriteCount={0}
        replyCount={0}
        shareCount={0}
        userId="test"
        userName="テストユーザー"
      />
      <PlatCell
        content="はじめてのplat"
        favoriteCount={0}
        replyCount={0}
        shareCount={0}
        userId="test"
        userName="テストユーザー"
      />
      <PlatCell
        content="はじめてのplat"
        favoriteCount={0}
        replyCount={0}
        shareCount={0}
        userId="test"
        userName="テストユーザー"
      />
      <PlatCell
        content="はじめてのplat"
        favoriteCount={0}
        replyCount={0}
        shareCount={0}
        userId="test"
        userName="テストユーザー"
      />
    </div>
  );
};
