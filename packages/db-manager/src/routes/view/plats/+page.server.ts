import { prisma } from '$lib/prisma.server';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async () => {
	const [plats, count] = await Promise.all([
		await prisma.plats.findMany({ orderBy: { created_at: 'desc' }, take: 10000 }),
		await prisma.plats.count()
	]);

	return { plats, count };
};
