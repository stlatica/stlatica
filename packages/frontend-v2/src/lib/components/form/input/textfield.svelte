<script lang="ts">
	import type { HTMLInputAttributes } from 'svelte/elements';

	type Props = {
		customProp?: boolean;
	} & HTMLInputAttributes;

	let props: Props = $props();

	// console.log({ ...props });
</script>

<input {...props} />

<style>
	input {
		/* material design の filled 風デザイン */
		background-color: inherit;
		border: 1px solid #fff;
		border-radius: var(--size-03);
		padding: 8px;
		width: 100%;
		color: var(--color-text);
	}
</style>
