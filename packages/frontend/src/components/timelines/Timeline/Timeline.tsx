import { Button } from "@mantine/core";
import { useVirtualizer } from "@tanstack/react-virtual";
import React, { useEffect } from "react";

import { PlatCellWithFetch } from "@/components/common/PlatCell/PlatCellWithFetch";
import { useTimeline } from "@/features/globalStore/TimelineStore";
import { useGetTimelineByQuery } from "@/openapi/stlaticaInternalApi";

import * as styles from "./Timeline.css";

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

  const [date] = React.useState(new Date().toISOString());

  const { data: tl } = useGetTimelineByQuery(
    { user_id, type: "home", to_date: date },
    {
      // 認証ヘッダー仮置き
      axios: {
        headers: {
          Authorization: "Bearer aaaa",
        },
      },
      // これだと時刻を変更して再取得が出来ないので要修正
      swr: { refreshInterval: 1000 },
    },
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
    <>
      <Button
        fullWidth={true}
        variant="outline"
        onClick={() => {
          virtualizer.scrollToIndex(0);
        }}
      >
        scroll to the top
      </Button>
      <div ref={parentRef} className={styles.container}>
        <div
          className={styles.scrollArea}
          // eslint-disable-next-line react/forbid-dom-props
          style={{
            height: virtualizer.getTotalSize(),
          }}
        >
          <div
            className={styles.positionHelper}
            // eslint-disable-next-line react/forbid-dom-props
            style={{
              transform: `translateY(${String(items[0]?.start ?? 0)}px)`,
            }}
          >
            {items.map((virtualRow) => {
              const { plat_id } = data[virtualRow.index];
              return (
                <div
                  key={virtualRow.key}
                  data-index={virtualRow.index}
                  ref={virtualizer.measureElement}
                  className={styles.item}
                >
                  <PlatCellWithFetch key={plat_id} id={plat_id} />
                </div>
              );
            })}
          </div>
        </div>
      </div>
    </>
  );
};
