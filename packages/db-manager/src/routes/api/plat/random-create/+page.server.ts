import { getTimelineByQuery } from '$lib/orval/stlaticaInternalApi';
import { prisma } from '$lib/prisma.server';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async () => {
	const users = await prisma.users.findMany({ orderBy: { preferred_user_id: 'asc' } });

	const timeline = await getTimelineByQuery({
		user_id: '01J8Q2FZM49SM8V4H288EFCBZF',
		type: 'home',
		to_date: new Date().toISOString()
	});

	return {
		users,
		timeline: timeline.data
	};
};
