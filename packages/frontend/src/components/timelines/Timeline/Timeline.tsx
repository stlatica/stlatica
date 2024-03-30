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
    // TODO: 90vh指定ではなく全高さを使うようにしたい
    <div className="h-[90vh] overflow-y-scroll">
      <div className="flex h-full flex-col gap-2">
        <PlatCell
          content="はじめてのplat"
          createdAt="2021-01-30T00:00:00+09:00"
          favoriteCount={0}
          replyCount={0}
          shareCount={0}
          userId="test"
          userName="テストユーザー"
        />
        <PlatCell
          content="はじめてのplat"
          createdAt="2021-01-30T00:00:00+09:00"
          favoriteCount={0}
          replyCount={0}
          shareCount={0}
          userId="test"
          userName="テストユーザー"
        />
        <PlatCell
          content="はじめてのplat"
          createdAt="2021-01-30T00:00:00+09:00"
          favoriteCount={0}
          replyCount={0}
          shareCount={0}
          userId="test"
          userName="テストユーザー"
        />
        <PlatCell
          content="はじめてのplat"
          createdAt="2021-01-30T00:00:00+09:00"
          favoriteCount={0}
          replyCount={0}
          shareCount={0}
          userId="test"
          userName="テストユーザー"
        />
        <PlatCell
          content="はじめてのplat"
          createdAt="2021-01-30T00:00:00+09:00"
          favoriteCount={0}
          replyCount={0}
          shareCount={0}
          userId="test"
          userName="テストユーザー"
        />
      </div>
    </div>
  );
};
