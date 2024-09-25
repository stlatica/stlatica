import { prisma } from '$lib/prisma.server';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async () => {
	const users = await prisma.users.findMany({ orderBy: { preferred_user_id: 'asc' } });

	return { users };
};
