/** @type {import('@ladle/react').UserConfig} */
export default {
  // ladleにbase指定していると meta.json が404になってしまう
  // base: "/stlatica/",
  outDir: "ladle-build",
  host: "0.0.0.0",
  viteConfig: `${process.cwd()}/.ladle/vite.ladle.config.ts`,
};
