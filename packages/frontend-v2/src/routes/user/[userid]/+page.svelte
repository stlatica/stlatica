<script>
	// import Plat from "$lib/components/timelines/plat.svelte";
	import Timeline from "$lib/components/timelines/timeline.svelte";
	import UserProfilePanel from "$lib/components/timelines/user-profile-panel.svelte";
</script>

<div class="page">
	<div class="panel">
		<UserProfilePanel></UserProfilePanel>
		<Timeline></Timeline>
	</div>
</div>

<style>
	.page {
		display: flex;
		justify-content: center;
		align-items: center;
	}

	.panel {
		max-width: 700px;
		border: solid 2px var(--gray-8);
	}
</style>
