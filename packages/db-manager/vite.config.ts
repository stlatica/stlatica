import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vitest/config';
import pluginPurgeCss from '@mojojoejo/vite-plugin-purgecss';

export default defineConfig({
	plugins: [sveltekit(), pluginPurgeCss()],
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	}
});
