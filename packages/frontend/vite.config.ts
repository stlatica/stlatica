import { defineConfig } from 'vitest/config';
import { sveltekit } from '@sveltejs/kit/vite';
import Icons from 'unplugin-icons/vite';
import { kitRoutes } from 'vite-plugin-kit-routes';

export default defineConfig({
	plugins: [
		sveltekit(),
		kitRoutes({
			generated_file_path: 'src/lib/routes/ROUTES.ts',
			format: 'object[path]'
		}),
		Icons({
			compiler: 'svelte'
		})
	],
	server: {
		host: '0.0.0.0'
	},
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	}
});
