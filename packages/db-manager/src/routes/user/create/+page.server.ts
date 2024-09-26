import type { PageServerLoad, Actions } from './$types';
import { prisma } from '$lib/prisma.server';
import { superValidate } from 'sveltekit-superforms';
import { valibot } from 'sveltekit-superforms/adapters';
import { fail } from '@sveltejs/kit';
import { FakeUser } from '$lib/faker/user';

import * as v from 'valibot';
import { generateUlid } from '$lib/faker/ulid';

const userSchema = v.object({
	user_id: v.pipe(v.string(), v.ulid()),
	preferred_user_id: v.string(),
	preferred_user_name: v.string(),
	is_public: v.boolean(),
	mail_address: v.string()
});

export const actions: Actions = {
	default: async (x) => {
		const form = await superValidate(x.request, valibot(userSchema));

		console.log(form);

		if (!form.valid) {
			return fail(400, { message: 'invalid form' + JSON.stringify(form.errors, null, 2) });
		}

		try {
			const r = await prisma.users.create({
				data: {
					...form.data,
					registered_at: new Date().valueOf(),
					created_at: new Date().valueOf(),
					updated_at: new Date().valueOf(),
					icon_image_id: generateUlid()
				}
			});

			return { message: `insert data: ${JSON.stringify(r, null, 2)}` };
		} catch (e) {
			console.error(e);
			return { message: String(e) };
		}
	}
};

export const load: PageServerLoad = async () => {
	// 初期データ返却(.svelte側でやるとFOUC発生する)
	const user = FakeUser();

	const form = await superValidate(user, valibot(userSchema));

	return { form };
};
