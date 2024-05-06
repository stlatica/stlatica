import { Stack } from "@mantine/core";
import React, { useEffect, useRef, useState } from "react";

import { IsNull, OnlyChildren } from "@/utilities/utilities";

type TimelineProps = {
  readonly elements: React.ReactElement[];
};

const HeightVisualizer: React.FC<OnlyChildren> = ({ children }) => {
  const ref = useRef<HTMLDivElement>(null);
  const [h, SetHeight] = useState(-1);

  useEffect(() => {
    const info = ref.current?.getBoundingClientRect();

    if (IsNull(info)) {
      return;
    }

    if (info.height === h) {
      return;
    }

    SetHeight(info.height);
  });

  return (
    <div ref={ref}>
      <div>height:{h}</div>
      {children}
    </div>
  );
};

/**
 *
 */
export const VirtualList: React.FC<TimelineProps> = ({ elements }) => {
  return (
    <Stack>
      {elements.map((x) => {
        return <HeightVisualizer>{x}</HeightVisualizer>;
      })}
    </Stack>
  );
};
