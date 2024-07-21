/* eslint-disable import/no-extraneous-dependencies */
import { PlaywrightTestConfig } from "@playwright/test";

export default {
  webServer: [
    {
      command: "pnpm run ladle serve",
      // process.env.E2E_ENV === "dev"
      //   ? "pnpm ladle serve"
      //   : "pnpm ladle build && pnpm ladle preview -p 61000",
      url: `http://localhost:61000`,
      reuseExistingServer: true,
    },
  ],
  reporter: [["html", { host: "0.0.0.0", outputFolder: "e2e/report" }]],
  use: {},
  projects: [
    {
      name: "visual-snapshot",
      testDir: "e2e/ladle-snapshots",
      snapshotDir: "e2e/ladle-snapshots",
      outputDir: "e2e/ladle-snapshots/output",
      fullyParallel: true,
    },
  ],
  // snapshotPathTemplate: "{testDir}/{testFilePath}-snapshots/{arg}-{projectName}{ext}",
} satisfies PlaywrightTestConfig;
