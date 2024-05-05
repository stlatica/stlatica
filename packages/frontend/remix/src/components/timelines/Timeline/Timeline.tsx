import { Stack } from "@mantine/core";
import React from "react";

import { PlatCell } from "@/components/common/PlatCell";

import { container } from "./Timeline.css";

type TimelineProps = {
  // children: React.ReactNode;
};

/**
 *
 */
export const Timeline: React.FC<TimelineProps> = () => {
  return (
    <Stack className={container}>
      {new Array(10).fill(0).map((_, i) => {
        return (
          <PlatCell
            key={i}
            content="はじめてのplat"
            favoriteCount={i * 10}
            replyCount={0}
            shareCount={i}
            userId="test"
            userName="テストユーザー"
          />
        );
      })}
    </Stack>
  );
};
