import { browser } from "$app/environment";

/**
 * クライアントサイドでのみ fetch を実行するためのヘルパー
 */
export const svelte_client_fetch = async <T>(func: () => T): Promise<T> => {
	if (!browser) {
		return await new Promise<T>(() => {});
	}

	return func();
};
