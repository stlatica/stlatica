import { Stack } from "@mantine/core";
import React, { useEffect } from "react";

import { PlatCellWithFetch } from "@/components/common/PlatCell/PlatCellWithFetch";
import { useTimeline } from "@/features/globalStore/TimelineStore";
import { useGetTimelineByQuery } from "@/openapi/stlaticaInternalApi";

import { container } from "./Timeline.css";

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
export const Timeline: React.FC<TimelineProps> = ({ url, user_id }) => {
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

  // // タイムラインAPIからTL取得していないので後でちゃんと置き換える
  // useSWR(
  //   url,
  //   async (u) => {
  //     const res = await fetcher(u);
  //     mutate(res);
  //     return res;
  //   },
  //   {
  //     refreshInterval: 1000,
  //   }
  // );

  return (
    <Stack className={container}>
      {data.map(({ plat_id }, i) => {
        return <PlatCellWithFetch key={i} id={plat_id} />;
      })}
    </Stack>
  );
};
