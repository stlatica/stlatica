/** @type {import('eslint').Linter.Config} */
module.exports = {
  root: true,
  parserOptions: {
    ecmaVersion: "latest",
    sourceType: "module",
    ecmaFeatures: {
      jsx: true,
    },
    project: ["./tsconfig.json"],
  },
  env: {
    browser: true,
    commonjs: true,
    es6: true,
  },

  // Base config
  extends: [
    "eslint:recommended",
    // import
    "plugin:import/recommended",
    "plugin:import/typescript",
    // typescript-eslintの設定解説→ https://zenn.dev/cybozu_frontend/articles/ts-eslint-v6-guide
    "plugin:@typescript-eslint/strict-type-checked",
    "plugin:promise/recommended",
    "plugin:tailwindcss/recommended",
    // prettierと競合するeslintの設定を除外する 末尾に入れておくこと
    "prettier",
  ],

  overrides: [
    // React
    {
      files: ["**/*.{js,jsx,ts,tsx}"],
      parserOptions: {
        project: `./tsconfig.json`,
      },
      plugins: [
        "react",
        "jsx-a11y",
        "import",
        "@typescript-eslint",
        "eslint-plugin-jsdoc",
        // @stylistic/eslint-plugin
        "@stylistic",
        "html",
        "promise",
        "unused-imports",
      ],
      extends: [
        "plugin:react/recommended",
        "plugin:react/jsx-runtime",
        "plugin:react-hooks/recommended",
        "plugin:jsx-a11y/recommended",
      ],
      settings: {
        react: {
          version: "detect",
        },
        formComponents: ["Form"],
        linkComponents: [
          { name: "Link", linkAttribute: "to" },
          { name: "NavLink", linkAttribute: "to" },
        ],
        "import/resolver": {
          typescript: {},
        },
      },
      rules: {
        "react/prop-types": "off", // react設定エリア
        "react/jsx-no-bind": [
          "error",
          {
            allowArrowFunctions: true,
          },
        ],
        // 赤字ではなく黄色で
        "react/self-closing-comp": "warn",
        // 赤字ではなく黄色で
        "react/prefer-read-only-props": "warn",
        // <div>Hoge</div> みたいなのをNGにしない
        "react/jsx-no-literals": "off",
        // JSXの深さ制限
        "react/jsx-max-depth": [
          "error",
          {
            max: 5,
          },
        ],
        // style要素の禁止
        "react/forbid-dom-props": [
          "error",
          {
            forbid: [
              {
                propName: "style",
                message:
                  "CSS styles is must written in css files. (If you need dynamic style, please disable this rule.)",
              },
            ],
          },
        ],
        // tailwind強制用: 同上 こちらはコンポーネント用
        "react/forbid-component-props": [
          "error",
          {
            forbid: [
              {
                propName: "style",
                message: "Use className (Must use tailwind css)",
              },
            ],
          },
        ],
        "react/jsx-no-useless-fragment": "off",
        // tsでは完全に不要なルール
        "react/require-default-props": "off",
        "react/prop-types": "off",
        // import React を不要にする
        "react/react-in-jsx-scope": "off",
        "react/jsx-props-no-spreading": "off",
      },
    },

    // Typescript
    {
      files: ["**/*.{ts,tsx}"],
      plugins: ["@typescript-eslint", "import"],
      parser: "@typescript-eslint/parser",
      settings: {
        "import/internal-regex": "^~/",
        "import/resolver": {
          node: {
            extensions: [".ts", ".tsx"],
          },
          typescript: {
            alwaysTryTypes: true,
          },
        },
      },
      extends: [
        "plugin:@typescript-eslint/recommended",
        "plugin:import/recommended",
        "plugin:import/typescript",
      ],
    },

    // Node
    {
      files: [".eslintrc.cjs"],
      env: {
        node: true,
      },
    },

    // .tsファイルではtsdocを必須に
    {
      files: ["*/**/*.ts"],
      rules: {
        "jsdoc/require-jsdoc": [
          "warn",
          {
            publicOnly: true,
            require: {
              ArrowFunctionExpression: true,
              ClassDeclaration: true,
              MethodDefinition: true,
            },
            checkConstructors: false,
          },
        ],
      },
    },

    // ストーリー, テストファイルでの制限緩和
    {
      files: ["*/**/*.stories.tsx", "*/**/*.test.tsx"],
      rules: {
        "react/no-multi-comp": "off",
        "import/no-extraneous-dependencies": "off",
        "react/forbid-dom-props": "off",
        "no-console": "off",
        "jsdoc/require-param": "off",
      },
    },
  ],

  settings: {
    "import/resolver": {
      node: {
        paths: ["src"],
        extensions: [".js", ".jsx", ".ts", ".tsx"],
      },
      typescript: {
        project: "./tsconfig.json",
      },
    },
  },

  rules: {
    // 1ファイル行数制限
    "max-lines": [
      "warn",
      {
        max: 250,
        // "skipComments": true,
        skipBlankLines: true,
      },
    ],
    // デフォルトだとtsxファイルでjsx書けないので許可
    "react/jsx-filename-extension": [
      "error",
      {
        extensions: [".jsx", ".tsx"],
      },
    ],
    "spaced-comment": [
      "warn",
      "always",
      {
        block: {
          balanced: true,
        },
      },
    ],
    "@typescript-eslint/no-use-before-define": [
      "error",
      {
        variables: true,
        functions: false,
        classes: true,
        typedefs: false,
      },
    ],
    "import/prefer-default-export": ["off"],
    // アロー関数を強制
    "func-style": ["error", "expression"],
    // アロー関数はカッコ必須
    "arrow-parens": [
      "error",
      "always",
      {
        requireForBlockBody: true,
      },
    ],
    // "arrow-body-style": "off",
    // typescriptでimport解決されるので不要.
    "node/no-missing-import": "off",
    "@typescript-eslint/no-empty-function": "off",
    "@typescript-eslint/explicit-module-boundary-types": "off",
    "react/function-component-definition": [
      "error",
      {
        namedComponents: "arrow-function",
        unnamedComponents: "arrow-function",
      },
    ],
    "no-use-before-define": "off",
    // console が本番に残っていると良くないので、マージまでにコメントアウトしておくこと
    "no-console": "warn",
    "@typescript-eslint/interface-name-prefix": ["off"],
    // interfaceを禁止, typeに統一
    "@typescript-eslint/consistent-type-definitions": ["error", "type"],
    // 使うべきシーンもあるので赤エラーは出さない. 必要なら随時無効化で対応を
    "no-await-in-loop": "warn",
    // ts, tsxのみ拡張子なし その他は必須とする
    // "import/extensions": [
    //   "error",
    //   "always",
    //   {
    //     ts: "never",
    //     tsx: "never",
    //     js: "never",
    //     jsx: "never",
    //   },
    // ],
    // 上に戻る相対パス禁止, @/ を使うこと
    "no-restricted-imports": [
      "error",
      {
        patterns: ["../"],
      },
    ],
    "@typescript-eslint/no-misused-promises": [
      "error",
      {
        checksVoidReturn: false,
      },
    ],
    "@typescript-eslint/no-floating-promises": "off",

    "arrow-body-style": ["error", "always"],
    "@stylistic/comma-dangle": ["error", "only-multiline"],
    "no-undefined": "off",
    "import/no-extraneous-dependencies": [
      "error",
      {
        devDependencies: ["**/*test.+(ts|tsx|js)"],
        peerDependencies: false,
      },
    ],
    // https://github.com/sweepline/eslint-plugin-unused-imports
    "@typescript-eslint/no-unused-vars": "off",
    "unused-imports/no-unused-imports": "warn",
    "unused-imports/no-unused-vars": [
      "warn",
      // { "vars": "all", "varsIgnorePattern": "^_", "args": "after-used", "argsIgnorePattern": "^_" }
    ],
    // import order
    // https://zenn.dev/knowledgework/articles/0994f518015c04?redirected=1#import%E3%81%AE%E8%87%AA%E5%8B%95%E6%95%B4%E5%88%97%EF%BC%88import%2Forder%EF%BC%89
    "import/order": [
      "warn",
      {
        groups: [
          "builtin",
          "external",
          "internal",
          ["parent", "sibling"],
          "object",
          "type",
          "index",
        ],
        "newlines-between": "always",
        pathGroupsExcludedImportTypes: ["builtin"],
        alphabetize: {
          order: "asc",
          caseInsensitive: true,
        },
      },
    ],
  },
  ignorePatterns: [
    "*.png",
    "*.jpg",
    "src/openapi/**/*",
    "**/*.config.js",
    "**/*.config.ts",
    "**/.eslintrc.*",
    "plop/**/*",
  ],
};
