"use client";

import { Tabs } from "@mantine/core";
import React from "react";

import { Timeline } from "@/components/timelines/Timeline";

type ProfileTabProps = {
  // children: React.ReactNode;
};

/**
 *
 */
export const ProfileTab: React.FC<ProfileTabProps> = () => {
  return (
    <div>
      <Tabs defaultValue="timeline" radius="xl" variant="default">
        <Tabs.List>
          <Tabs.Tab value="timeline">Timeline</Tabs.Tab>
          <Tabs.Tab value="reply">Reply</Tabs.Tab>
          <Tabs.Tab value="media">Media</Tabs.Tab>
          <Tabs.Tab value="favorite">Favorite</Tabs.Tab>
        </Tabs.List>

        <Tabs.Panel value="timeline">
          <div className="mt-2">
            <Timeline />
          </div>
        </Tabs.Panel>
        <Tabs.Panel value="reply">
          <div className="mt-2">
            <Timeline />
          </div>
        </Tabs.Panel>
        <Tabs.Panel value="media">
          <div className="mt-2">
            <Timeline />
          </div>
        </Tabs.Panel>
        <Tabs.Panel value="favorite">
          <div className="mt-2">
            <Timeline />
          </div>
        </Tabs.Panel>
      </Tabs>
    </div>
  );
};
