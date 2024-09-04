<script lang="ts">
	import type { ButtonProps } from './button';

	type Props = ButtonProps;

	let { children, onclick, fullwidth, variant = 'outlined', ...others }: Props = $props();

	// console.log(fullwidth);
	const classNames = [fullwidth ? 'fullwidth' : '', variant].join(' ');
	// console.log(classNames);
</script>

<button class={`button-base ${classNames}`} {onclick} {...others}>
	{@render children?.()}
</button>

<style>
	button {
		background-color: var(--blue-10);
		padding: var(--size-4);
		border-radius: var(--rounded-lg);
		text-align: center;
		text-decoration: none;
		display: inline-block;
		cursor: pointer;
		transition: all var(--transition-duration-ui) ease-out;
	}
	button:active {
		background-color: var(--blue-11);
	}
	button:active {
		transform: scale(1.03) translateY(2px);
		box-shadow: none;
	}

	button:hover {
		transform: scale(1.03);
	}

	.fullwidth {
		width: 100%;
	}

	/* outlined */

	.outlined {
		background-color: inherit;
		border: var(--border-width-2) var(--blue-7);
		color: var(--blue-11);
	}
	.outlined:active {
		/* 透明度90%の灰色の背景色を hsl で設定 */
		background-color: hsla(0, 0%, 90%, 0.1);
	}

	/* contained */

	.contained {
		border: none;
		color: hsl(0, 0%, 95%);
		background-color: var(--blue-10);
	}
	.button-base:active {
		background-color: var(--blue-11);
	}
</style>
