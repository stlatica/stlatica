import { Hono } from "hono";
import { describeRoute } from "hono-openapi";
import { resolver } from "hono-openapi/valibot";
import * as v from "valibot";
import { openAPISpecs } from "hono-openapi";
import { apiReference } from "@scalar/hono-api-reference";

const app = new Hono();

const responseSchema = v.string();

app.get(
  "bff/hono",

  describeRoute({
    description: "Say hello to the user",
    responses: {
      200: {
        description: "Successful response",
        content: {
          "text/plain": { schema: resolver(responseSchema) },
        },
      },
    },
  }),
  (c) => c.text("Hono!"),
);

const baseEndpoint = "/bff";

app.get(
  "bff/openapi",
  openAPISpecs(app, {
    documentation: {
      info: { title: "Hono API", version: "1.0.0", description: "Greeting API" },
      servers: [{ url: "http://localhost:3000", description: "Local Server" }],
    },
  }),
);

app.get(
  "bff/docs",
  apiReference({
    theme: "saturn",
    spec: { url: "/bff/openapi" },
  }),
);

export const honoApp = new Hono({ strict: false }).route("/", app);
export type AppType = typeof honoApp;

export default honoApp;
