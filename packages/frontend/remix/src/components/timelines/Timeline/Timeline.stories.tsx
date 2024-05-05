import { GenerateDummyPlat } from "@/features/test/GenerateDummyPlat";

import { Timeline } from "./Timeline";

import type { Plat } from "@/openapi/stlaticaInternalApi.schemas";
import type { Story, StoryDefault } from "@ladle/react";

export default {
  title: "timelines/Timeline",
} satisfies StoryDefault;

export const Story1: Story = () => {
  const data = new Array(20).fill(0).map(() => {
    return GenerateDummyPlat();
  });
  const fetcher = () => {
    return new Promise<Plat[]>((resolve) => {
      resolve(data);
    });
  };

  return <Timeline url="story/timeline/story1" fetcher={fetcher} />;
};
Story1.storyName = "1.基本";

export const Story2: Story = () => {
  const fetcher = () => {
    return new Promise<Plat[]>((resolve) => {
      resolve([GenerateDummyPlat(), GenerateDummyPlat(), GenerateDummyPlat()]);
    });
  };
  return <Timeline url="story/timeline/story2" fetcher={fetcher} />;
};
Story2.storyName = "2.仮動作確認";
