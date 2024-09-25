import { Hono } from 'hono';
// import { vValidator } from '@hono/valibot-validator';
import { fakerJA } from '@faker-js/faker';
import { ulid } from 'ulidx';

const router = new Hono()
	.get('/hello', async (c) => {
		// await new Promise((resolve) => {
		// 	setTimeout(() => {
		// 		resolve(0);
		// 	}, 10000);
		// });
		return c.json({ message: 'hello api', time: new Date(), method: 'GET' });
	})
	.post('/hello', async (c) => {
		return c.json({ message: 'hello api', time: new Date(), method: 'POST' });
	})
	.delete('/hello', async (c) => {
		return c.json({ message: 'hello api', time: new Date(), method: 'DELETE' });
	});

const router2 = new Hono().get('/faker/user', async (c) => {
	return c.json({
		user_id: ulid(),
		preferred_user_id: fakerJA.string.alphanumeric({ length: { min: 8, max: 16 } }),
		preferred_user_name: fakerJA.person.fullName(),
		is_public: true,
		mail_address: fakerJA.internet.email()
	});
});

// export const app = new Hono().route("/api", router);

export const honoApp = new Hono({ strict: false }).route('/api', router).route('/api', router2);

export type AppType = typeof honoApp;
