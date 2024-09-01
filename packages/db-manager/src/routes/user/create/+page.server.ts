import type { PageServerLoad, Actions } from './$types';
import { honoClient } from '$lib/hono/honoClient';
import { prisma } from '$lib/prisma.server';

export const actions: Actions = {
	default: async (x) => {
		const form = await x.request.formData();

		console.log(form.get('is_public'));

		try {
			const r = await prisma.users.create({
				data: {
					// ...formObject,
					user_id: String(form.get('user_id')),
					preferred_user_id: String(form.get('preferred_user_id')),
					preferred_user_name: String(form.get('preferred_user_name')),
					mail_address: String(form.get('mail_address')),
					is_public: form.get('is_public') === 'on',
					registered_at: new Date().valueOf(),
					created_at: new Date().valueOf(),
					updated_at: new Date().valueOf()
				}
			});

			// TODO log the user in
			return { success: true, data: r };
		} catch (e) {
			console.error(e);
			return { success: false, message: String(e) };
		}
	}
};

export const load: PageServerLoad = async () => {
	// 初期データ返却(.svelte側でやるとFOUC発生する)
	const user = await (await honoClient.api.faker.user.$get()).json();

	return {
		props: { ...user }
	};
};
