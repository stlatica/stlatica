<script lang="ts">
	import '$lib/styles/tw-variables/_index.css';
	import '$lib/styles/theme.css';
	import '$lib/styles/global.css';
	import '$lib/styles/utilities.css';
	import { PAGES } from '$lib/routes/ROUTES';
	import { QueryClient, QueryClientProvider } from '@tanstack/svelte-query';

	const queryClient = new QueryClient();
</script>

<div class="base">
	<div class="slot">
		<QueryClientProvider client={queryClient}>
			<slot></slot>
		</QueryClientProvider>
	</div>
</div>
<footer>
	<a href={PAGES['/']}> debug用: top画面へ </a>
</footer>

<style>
	.base {
		display: flex;

		height: 100%;
		/* background-color: blue; */
		flex-direction: row;
	}

	.slot {
		flex: 1;
		overflow: auto;
	}

	footer {
		/* 画面下部に固定する */
		position: fixed;
		bottom: 0;
		width: 100%;
	}
</style>
