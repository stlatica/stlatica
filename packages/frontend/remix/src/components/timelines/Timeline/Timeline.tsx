import { useVirtualizer } from "@tanstack/react-virtual";
import React, { useEffect } from "react";

import { PlatCellWithFetch } from "@/components/common/PlatCell/PlatCellWithFetch";
import { useTimeline } from "@/features/globalStore/TimelineStore";
import { useGetTimelineByQuery } from "@/openapi/stlaticaInternalApi";

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

  const parentRef = React.useRef<HTMLDivElement>(null);

  const count = data.length;
  const virtualizer = useVirtualizer({
    count,
    getScrollElement: () => {
      return parentRef.current;
    },
    estimateSize: () => {
      return 45;
    },
  });

  const items = virtualizer.getVirtualItems();

  return (
    <div
      ref={parentRef}
      style={{
        height: "100%",
        overflowY: "auto",
        contain: "strict",
      }}
    >
      <div
        style={{
          height: virtualizer.getTotalSize(),
          width: "100%",
          position: "relative",
        }}
      >
        <div
          style={{
            position: "absolute",
            top: 0,
            left: 0,
            width: "100%",
            transform: `translateY(${items[0]?.start ?? 0}px)`,
          }}
        >
          {items.map((virtualRow) => {
            const { plat_id } = data[virtualRow.index];
            return (
              <div
                key={virtualRow.key}
                data-index={virtualRow.index}
                ref={virtualizer.measureElement}
                className={virtualRow.index % 2 ? "ListItemOdd" : "ListItemEven"}
              >
                <PlatCellWithFetch key={plat_id} id={plat_id} />

                {/* <div style={{ padding: "10px 0" }}>
                    <div>Row {virtualRow.index}</div>
                    <div>{sentences[virtualRow.index]}</div>
                  </div> */}
              </div>
            );
          })}
        </div>
      </div>
    </div>
  );

  // return (
  //   <Stack className={container}>
  //     {data.map(({ plat_id }, i) => {
  //       return <PlatCellWithFetch key={i} id={plat_id} />;
  //     })}
  //   </Stack>
  // );
};
