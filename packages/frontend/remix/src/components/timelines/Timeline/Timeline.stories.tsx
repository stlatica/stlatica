import { GenerateDummyPlat } from "@/features/test/GenerateDummyPlat";

import { Timeline } from "./Timeline";

import type { Plat } from "@/openapi/stlaticaInternalApi.schemas";
import type { Story, StoryDefault } from "@ladle/react";

export default {
  title: "timelines/Timeline",
} satisfies StoryDefault;

const fetcher = () => {
  return new Promise<Plat[]>((resolve) => {
    resolve([GenerateDummyPlat(), GenerateDummyPlat(), GenerateDummyPlat()]);
  });
};

export const Story1: Story = () => {
  return <Timeline url="story/timeline" fetcher={fetcher} />;
};
