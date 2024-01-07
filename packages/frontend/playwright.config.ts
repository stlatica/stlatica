/* eslint-disable import/no-extraneous-dependencies */
import { PlaywrightTestConfig } from "@playwright/test";

export default {
  webServer: {
    command: "pnpm ladle build && pnpm ladle preview -p 61000",
    // command:
    //   process.env.TYPE === "dev" ? "pnpm ladle serve" : "pnpm ladle build && pnpm ladle preview",
    url: `http://localhost:61000/stlatica`,
  },
  reporter: [["html", { outputFolder: "playwright/report" }]],
  projects: [
    {
      name: "visual-snapshot",
      testDir: "playwright",
      snapshotDir: "playwright/snapshots",
      outputDir: "playwright/snapshots/output",
    },
  ],
} satisfies PlaywrightTestConfig;
