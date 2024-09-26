<script lang="ts">
	// let result = getTimelineByQuery({
	// 	user_id: '01J6YSX5HRSEDS46XGTD57P3D6',
	// 	type: 'home',
	// 	to_date: new Date().toISOString()
	// });

	export let data;

	// $: console.log(data);

	// $: console.log(result);
</script>

<p>デバッグ用: plat取得</p>

<pre>
{JSON.stringify(data.timeline, null, 2)}
</pre>

<!-- 
{#await result}
	<p>Loading...</p>
{:then users}
	<pre>
  {JSON.stringify(users.data, null, 2)}
</pre>
{:catch error}
	<p>{error.message}</p>
{/await} -->
