import type { Snippet } from 'svelte';
import type { HTMLButtonAttributes } from 'svelte/elements';

export type ButtonProps = {
	/**
	 * html button type
	 */
	type: HTMLButtonAttributes['type'];
	children?: Snippet;
	onclick?: () => void;
	/**
	 * width: 100% を適用する
	 */
	fullwidth?: boolean;
	/**
	 * ボタン種別
	 */
	variant?: 'outlined' | 'contained';
};
