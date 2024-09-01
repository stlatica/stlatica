import { honoApp } from '$lib/hono/hono.server';

export const GET = async (event) => {
	return await honoApp.fetch(event.request);
};

export const POST = GET;
export const PUT = GET;
