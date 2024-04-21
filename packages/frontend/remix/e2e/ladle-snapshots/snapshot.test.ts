import { test, expect } from "@playwright/test";
// we can't create tests asynchronously, thus using the sync-fetch lib
import fetch from "sync-fetch";

// URL where Ladle is served
const url = "http://localhost:61000";

// TODO: sync-fetchを使わずに実行 https://github.com/stlatica/stlatica/issues/261
// fetch Ladle's meta file
// https://ladle.dev/docs/meta
// eslint-disable-next-line @typescript-eslint/no-unsafe-assignment, @typescript-eslint/no-unsafe-member-access
const stories = fetch(`${url}/meta.json`).json().stories;

// iterate through stories
// eslint-disable-next-line @typescript-eslint/no-unsafe-argument
Object.keys(stories).forEach((storyKey) => {
  // create a test for each story
  test(`snapshot test - ${storyKey}`, async ({ page }) => {
    // skip stories with `meta.skip` set to true
    // test.skip(stories[storyKey].meta.skip, "meta.skip is true");
    // navigate to the story
    await page.goto(`${url}/?story=${storyKey}&mode=preview`);
    // stories are code-splitted, wait for them to be loaded
    await page.waitForSelector("[data-storyloaded]");
    // take a screenshot and compare it with the baseline
    await expect(page).toHaveScreenshot(`${storyKey}.png`, {
      // CIとの環境差分で落ちる場合はここを調整
      // maxDiffPixels: 100,
      fullPage: true,
    });
  });
});
