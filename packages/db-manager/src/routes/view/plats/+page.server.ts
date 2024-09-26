import { prisma } from '$lib/prisma.server';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async () => {
	const plats = await prisma.plats.findMany({ orderBy: { created_at: 'desc' } });

	return { plats };
};
