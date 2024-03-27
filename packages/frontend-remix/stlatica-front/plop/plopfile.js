module.exports = function (
  /** @type {import('plop').NodePlopAPI} */
  plop
) {
  plop.setGenerator("component", {
    description: "Create a component",
    prompts: [
      {
        type: "list",
        name: "category",
        message: "カテゴリを選んでください",
        choices: [
          { name: "button", value: "button" },
          { name: "input", value: "input" },
          { name: "dialog", value: "dialog" },
          { name: "common", value: "common" },
          { name: "block", value: "block" },
          { name: "timelines", value: "timelines" },
          { name: "others", value: "others" },
        ],
      },
      {
        type: "input",
        name: "name",
        message: "コンポーネント名を入力してください",
      },
    ],
    actions: (data) => {
      const path = `../src/components/${data?.category}/`;

      const actions = [
        {
          type: "add",
          path: path + "{{pascalCase name}}/index.ts",
          templateFile: "components/index.ts.hbs",
        },
        {
          type: "add",
          path: path + `{{pascalCase name}}/{{pascalCase name}}.stories.tsx`,
          templateFile: `components/story.tsx.hbs`,
        },
        {
          type: "add",
          path: path + `{{pascalCase name}}/{{pascalCase name}}.tsx`,
          templateFile: `components/component.tsx.hbs`,
        },
      ];
      return actions;
    },
  });
};
