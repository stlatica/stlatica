import { defineConfig } from 'orval';

export default defineConfig({
	stlatica_localhost: {
		output: {
			mode: 'split',
			target: 'src/lib/orval/client',
			// schemas: '../shared/openapi/build/openapi-bundled-internal.yaml',
			client: 'svelte-query',
			// mock: true,
			baseUrl: 'http://localhost:8080'
		},
		input: {
			target: '../shared/openapi/build/openapi-bundled-internal.yaml'
		}
	},
	stlatica_docker: {
		output: {
			mode: 'split',
			target: 'src/lib/orval/docker/',
			// schemas: '../shared/openapi/build/openapi-bundled-internal.yaml',
			client: 'svelte-query',
			// mock: true,
			baseUrl: 'http://stlatica_server:8080'
		},
		input: {
			target: '../shared/openapi/build/openapi-bundled-internal.yaml'
		}
	}
});
