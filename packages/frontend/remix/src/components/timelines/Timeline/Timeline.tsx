import { Stack } from "@mantine/core";
import React, { useEffect } from "react";

<<<<<<< HEAD
import { PlatCell } from "@/components/common/PlatCell";
import { UserIconSampleImage } from "@/components/common/UserIcon/UserIconSampleImage";
=======
import { PlatCellWithFetch } from "@/components/common/PlatCell/PlatCellWithFetch";
import { useTimeline } from "@/features/globalStore/TimelineStore";
import { useGetTimelineByQuery } from "@/openapi/stlaticaInternalApi";

import { container } from "./Timeline.css";
>>>>>>> main

type TimelineProps = {
  /**
   * cache key
   */
  readonly url: string;
  readonly user_id: string;
};

/**
 *
 */
export const ProfileTimeline: React.FC<TimelineProps> = ({ url, user_id }) => {
  const { data, mutate } = useTimeline(url);

  const { data: tl } = useGetTimelineByQuery(
    { user_id, type: "home" },
    {
      // 認証ヘッダー仮置き
      axios: {
        headers: {
          Authorization: "Bearer aaaa",
        },
      },
      swr: { refreshInterval: 1000 },
    }
  );

  useEffect(() => {
    if (tl === undefined) {
      return;
    }
    mutate(tl.data);
  }, [mutate, tl]);

  return (
<<<<<<< HEAD
    // TODO: 90vh指定ではなく全高さを使うようにしたい
    <div className="h-[90vh] overflow-y-scroll">
      <div className="flex h-full flex-col gap-2">
        <PlatCell
          content="はじめてのplat"
          favoriteCount={0}
          replyCount={0}
          shareCount={0}
          userIcon={UserIconSampleImage}
          userId="test"
          userName="テストユーザー"
        />
        <PlatCell
          content="はじめてのplat"
          favoriteCount={0}
          replyCount={0}
          shareCount={0}
          userIcon={UserIconSampleImage}
          userId="test"
          userName="テストユーザー"
        />
        <PlatCell
          content="はじめてのplat"
          favoriteCount={0}
          replyCount={0}
          shareCount={0}
          userIcon={UserIconSampleImage}
          userId="test"
          userName="テストユーザー"
        />
        <PlatCell
          content="はじめてのplat"
          favoriteCount={0}
          replyCount={0}
          shareCount={0}
          userIcon={UserIconSampleImage}
          userId="test"
          userName="テストユーザー"
        />
        <PlatCell
          content="はじめてのplat"
          favoriteCount={0}
          replyCount={0}
          shareCount={0}
          userIcon={UserIconSampleImage}
          userId="test"
          userName="テストユーザー"
        />
      </div>
    </div>
=======
    <Stack className={container}>
      {data.map(({ plat_id }, i) => {
        return <PlatCellWithFetch key={i} id={plat_id} />;
      })}
    </Stack>
>>>>>>> main
  );
};
