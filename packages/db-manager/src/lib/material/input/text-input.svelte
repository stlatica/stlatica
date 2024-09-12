<script lang="ts">
	import type { HTMLInputAttributes } from 'svelte/elements';

	type Props = {
		customProp?: boolean;
	} & HTMLInputAttributes;

	let { value = $bindable(''), ...props }: Props = $props();
</script>

<input {...props} bind:value />

<style>
	input {
		/* material design の filled 風デザイン */
		background-color: inherit;
		border: 1px solid var(--gray-11);
		border-radius: var(--size-3);
		padding: 8px;
		/* width: 100%; */
		color: var(--color-text);
		flex: 1;
	}
</style>
