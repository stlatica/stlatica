<script lang="ts">
	import { createGetTimelineByQuery } from '$lib/orval/stlaticaInternalApi';
	import Plat from './plat.svelte';

	const query = createGetTimelineByQuery({
		user_id: '01J8Q2G2NK6KRWN85QPTD73D4M',
		type: 'home',
		to_date: new Date().toISOString()
	});

	console.log($query.data);
</script>

{#if $query.isLoading}
	<p>Loading...</p>
{:else if $query.isError}
	<p>Error: {$query.error.message}</p>
{:else}
	{#each $query.data.data as data}
		<Plat username={data.user_id} userid={data.user_id} content={data.content}></Plat>
	{/each}
{/if}
