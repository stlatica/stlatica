<script lang="ts">
	import Button from "$lib/components/material/button/button.svelte";

	type Props = {
		isFollow: boolean;
		onclick: () => Promise<void>;
	};

	let { isFollow, ...others }: Props = $props();
</script>

<!-- <div> -->
{#if isFollow}
	<Button type="button" variant="outlined" {...others}>フォロー解除</Button>
{:else}
	<Button type="button" variant="contained" {...others}>フォロー</Button>
{/if}

<!-- </div> -->

<style>
	/* div {
		display: inline-block;
		width: 8em;
		display: flex;
		justify-content: center;
	} */
</style>
