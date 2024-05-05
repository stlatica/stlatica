import { Stack } from "@mantine/core";
import React from "react";
import useSWR from "swr";

import { PlatCellWithFetch } from "@/components/common/PlatCell/PlatCellWithFetch";
import { useTimeline } from "@/features/globalStore/TimelineStore";
import { Plat } from "@/openapi/stlaticaInternalApi.schemas";

import { container } from "./Timeline.css";

type TimelineProps = {
  /**
   * cache key
   */
  readonly url: string;
  readonly fetcher: (url: string) => Promise<Plat[]>;
};

/**
 *
 */
export const Timeline: React.FC<TimelineProps> = ({ url, fetcher }) => {
  const { data, mutate } = useTimeline(url);

  useSWR(
    url,
    async (u) => {
      const res = await fetcher(u);
      mutate(res);
      return res;
    },
    {
      refreshInterval: 1000,
    }
  );

  return (
    <Stack className={container}>
      {data.map(({ plat_id }, i) => {
        return <PlatCellWithFetch key={i} id={plat_id} />;
      })}
    </Stack>
  );
};
