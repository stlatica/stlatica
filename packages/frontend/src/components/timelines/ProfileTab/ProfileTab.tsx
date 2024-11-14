import { Tabs } from "@mantine/core";
import { useParams } from "@remix-run/react";
import type React from "react";

import { useNavigateSafe } from "@/features/acl/useNavigate";

// type ProfileTabProps = {
//   // children: React.ReactNode;
// };

/**
 * Profile 表示用のタブ
 */
export const ProfileTab: React.FC = () => {
  const navigate = useNavigateSafe();
  const { tabValue } = useParams();

  return (
    <div>
      <Tabs
        defaultValue="timeline"
        radius="xl"
        variant="default"
        value={tabValue}
        onChange={(value) => {
          if (value === null) {
            return;
          }
          navigate(value);
        }}
      >
        <Tabs.List>
          <Tabs.Tab value="timeline">Timeline</Tabs.Tab>
          <Tabs.Tab value="media">Media</Tabs.Tab>
          <Tabs.Tab value="favorites">Favorites</Tabs.Tab>
          <Tabs.Tab value="follows">Follows</Tabs.Tab>
          <Tabs.Tab value="followers">Followers</Tabs.Tab>
        </Tabs.List>
      </Tabs>
    </div>
  );
};
