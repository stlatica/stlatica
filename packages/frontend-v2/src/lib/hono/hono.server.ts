import { Hono } from "hono";
import { vValidator } from "@hono/valibot-validator";
import * as v from "valibot";

const router = new Hono()
	.get("/hello", async (c) => {
		// await new Promise((resolve) => {
		// 	setTimeout(() => {
		// 		resolve(0);
		// 	}, 10000);
		// });
		return c.json({ message: "hello api", time: new Date(), method: "GET" });
	})
	.post("/hello", async (c) => {
		return c.json({ message: "hello api", time: new Date(), method: "POST" });
	})
	.delete("/hello", async (c) => {
		return c.json({ message: "hello api", time: new Date(), method: "DELETE" });
	});

const router2 = new Hono().post(
	"/validate",
	vValidator(
		"json",
		v.object({
			id: v.string()
		})
	),
	async (c) => {
		const req = c.req.valid("json");
		console.log(req);
		return c.json({ message: req, method: "POST" });
	}
);

// export const app = new Hono().route("/api", router);

const baseEndpoint = "/api";

export const honoApp = new Hono({ strict: false })
	.route(baseEndpoint, router)
	.route(baseEndpoint, router2);

export type AppType = typeof honoApp;
