/* eslint-disable import/no-extraneous-dependencies */
import { PlaywrightTestConfig } from "@playwright/test";

export default {
  webServer: [
    {
      command:
        process.env.E2E_ENV === "dev"
          ? "pnpm ladle serve"
          : "pnpm ladle build && pnpm ladle preview -p 61000",
      url: `http://localhost:61000`,
    },
  ],
  reporter: [["html", { outputFolder: "playwright/report" }]],
  projects: [
    {
      name: "visual-snapshot",
      testDir: "playwright/snapshots",
      snapshotDir: "playwright/snapshots",
      outputDir: "playwright/snapshots/output",
    },
  ],
  snapshotPathTemplate: "{testDir}/{testFilePath}-snapshots/{arg}-{projectName}{ext}",
} satisfies PlaywrightTestConfig;
