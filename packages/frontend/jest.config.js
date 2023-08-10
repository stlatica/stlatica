/** @type {import('ts-jest/dist/types').InitialOptionsTsJest} */
module.exports = {
  verbose: true,
  testEnvironment: "node",
  transform: {
    // nextjsがswcらしいので合わせる意図でesbuildではなくswc利用
    "^.+\\.tsx?$": ["@swc/jest"],
  },
  // transform: {
  //   "^.+\\.(ts|tsx)$": [
  //     "esbuild-jest",
  //     {
  //       sourcemap: true,
  //     },
  //   ],
  // },
  // transformIgnorePatterns: ["node_modules/"],

  // testPathIgnorePatterns: ["src/components/**/*.stories.*"],

  // testMatchはデフォルト設定で十分機能する
  // testMatch: ["**/*test.+(ts|tsx|js)"],

  collectCoverageFrom: [
    // テスト未記載の.tsファイルもカバレッジで列挙する
    "src/**/*.+(ts|tsx|js)",
    // ストーリーファイルは除外
    "!src/**/*.stories.+(ts|tsx|js)",
    // Pageファイルも除外
    "!src/app/**/page.tsx",
    // グローバル系を除外
    "!src/app/**/layout.tsx",
  ],

  // collectCoverageFromを指定した場合、テスト書かないファイルが全部列挙されてしまい邪魔な時は除外する設定
  // coveragePathIgnorePatterns: ["src/components/"],
};
