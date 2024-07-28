import { MantineProvider } from "@mantine/core";
import { SWRConfig } from "swr";

import { OnlyChildren } from "./utilities/utilities";

export const Providers: React.FC<OnlyChildren> = ({ children }) => {
  return (
    <SWRConfig
      value={{
        // https://swr.vercel.app/ja/docs/revalidation
        // revalidateIfStale: false,
        revalidateOnFocus: false,
        revalidateOnReconnect: false,
      }}
    >
      <MantineProvider defaultColorScheme="dark">{children}</MantineProvider>
    </SWRConfig>
  );
};
