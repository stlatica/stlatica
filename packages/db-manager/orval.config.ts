import { defineConfig } from 'orval';

export default defineConfig({
	petstore: {
		output: {
			mode: 'split',
			target: 'src/lib/orval/',
			// schemas: '../shared/openapi/build/openapi-bundled-internal.yaml',
			client: 'svelte-query',
			// mock: true,
			baseUrl: 'http://localhost:8080'
		},
		input: {
			target: '../shared/openapi/build/openapi-bundled-internal.yaml'
		}
	}
});
