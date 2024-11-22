<script>
	import Plat from "./plat.svelte";

	const data = new Array(100).fill(0);
</script>

{#each data as _, i}
	<Plat
		username="TestUser"
		userid="sample"
		content="sample plat sampleeeeeeeeeeeeeeeeeeeeee eeeeeeeeeeeeeeeeeeeeeeeeeeeeee eeeeeeeeeeeeeeeeeeeeあいうえおあいうえおあいうえおあいうえおあいうえおあいうえおあいうえお"
	></Plat>
{/each}
