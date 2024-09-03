<script lang="ts">
	import { getTimelineByQuery } from '$lib/orval/stlaticaInternalApi';

	let result = getTimelineByQuery({
		user_id: '0BsARp8vHrgd',
		type: 'home',
		to_date: new Date().toISOString()
	});

	$: console.log(result);
</script>

{#await result}
	<p>Loading...</p>
{:then users}
	<pre>
  {JSON.stringify(users.data, null, 2)}
</pre>
{:catch error}
	<p>{error.message}</p>
{/await}
