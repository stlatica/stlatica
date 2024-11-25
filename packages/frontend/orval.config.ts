import { defineConfig } from 'orval';

export default defineConfig({
	api: {
		output: {
			// .msw を別で生成するため
			mode: 'split',
			target: 'src/lib/orval/',
			client: 'svelte-query',
			mock: true,
			// useTypeOverInterfaces: true,
			baseUrl: 'http://localhost:8080'
		},
		input: {
			target: '../shared/openapi/build/openapi-bundled-internal.yaml'
		}
	}
});
