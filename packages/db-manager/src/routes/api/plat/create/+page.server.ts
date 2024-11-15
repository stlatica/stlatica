import { postPlat } from '$lib/orval/docker/stlaticaInternalApi';
import { getTimelineByQuery } from '$lib/orval/docker/stlaticaInternalApi';
import { prisma } from '$lib/prisma.server';
import { fail, superValidate } from 'sveltekit-superforms';
import { valibot } from 'sveltekit-superforms/adapters';
import type { PageServerLoad } from './$types';
import * as v from 'valibot';
import { fakerJA } from '@faker-js/faker';
import type { Actions } from '@sveltejs/kit';

const schema = v.object({
	user_id: v.pipe(v.string(), v.ulid()),
	content: v.string()
});

export const load: PageServerLoad = async () => {
	const users = await prisma.users.findMany({ orderBy: { preferred_user_id: 'asc' } });

	// users の中からランダムなひとつをピックアップする
	const randomIndex = Math.floor(Math.random() * users.length);
	const pickupUser = users[randomIndex];

	const timeline = await getTimelineByQuery({
		user_id: '01J8Q2FZM49SM8V4H288EFCBZF',
		type: 'home',
		to_date: new Date().toISOString()
	});
	const form = await superValidate(
		{ user_id: pickupUser.user_id, content: fakerJA.lorem.paragraph() },
		valibot(schema)
	);

	return {
		form,
		pickupUser,
		timeline: timeline.data
	};
};

export const actions: Actions = {
	default: async ({ request }) => {
		const form = await superValidate(request, valibot(schema));

		if (!form.valid) {
			return fail(400, { message: 'invalid form' + JSON.stringify(form.errors, null, 2) });
		}
		const r = await postPlat({
			user_id: form.data.user_id,
			content: form.data.content
		});

		console.log(r);

		return { success: true };
	}
};
