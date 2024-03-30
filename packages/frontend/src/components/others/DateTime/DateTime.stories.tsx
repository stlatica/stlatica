import React from "react";

import { DateTime } from "./DateTime";

import type { Story, StoryDefault } from "@ladle/react";

export default {
  title: "components/others/DateTime",
} satisfies StoryDefault;

export const Story1: Story = () => {
  const formatDateLocale = (date: Date): string => {
    return date.toLocaleString();
  };
  return <DateTime date="2021-01-30T00:00:00+09:00" formatDate={formatDateLocale} />;
};
Story1.storyName = "時刻と日付";

export const Story2: Story = () => {
  const formatDateLocale = (date: Date): string => {
    return date.toLocaleDateString();
  };
  return <DateTime date="2021-01-30T00:00:00+09:00" formatDate={formatDateLocale} />;
};
Story2.storyName = "日付のみ";

export const Story3: Story = () => {
  const formatDateLocale = (date: Date): string => {
    return date.toLocaleTimeString();
  };
  return <DateTime date="2021-01-30T00:00:00+09:00" formatDate={formatDateLocale} />;
};
Story3.storyName = "時刻のみ";
