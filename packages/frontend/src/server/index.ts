import { Hono } from "hono";

const app = new Hono();

app.get("/hono", (c) => c.text("Hono!"));

export default app;
