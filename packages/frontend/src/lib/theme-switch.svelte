<script lang="ts">
	let isDark: boolean | undefined = undefined;
	import { browser } from "$app/environment";

	if (browser) {
		isDark = document.documentElement.classList.contains("dark");
	}

	const onClick = () => {
		if (document.documentElement.classList.contains("dark")) {
			document.documentElement.classList.remove("dark");
			isDark = false;
		} else {
			document.documentElement.classList.add("dark");
			isDark = true;
		}
		localStorage.setItem("theme", isDark ? "dark" : "light");
	};

	// https://icon-sets.iconify.design/ic/?category=General
	import Moon from "virtual:icons/ic/outline-mode-night";
	import Sun from "virtual:icons/ic/outline-wb-sunny";
</script>

<button class="button" onclick={onClick}>
	<!-- isDark が true なら Moon, false なら Sun を表示する -->
	{#if isDark === undefined}{:else if isDark === true}
		<Sun></Sun>
	{:else}
		<Moon></Moon>
	{/if}
</button>

<style>
	button {
		/* background-color: var(--blue-9); */
		color: var(--gray-12);
		background-color: inherit;
		font-size: 1.2rem;
		border-color: var(--gray-8);
		padding: var(--size-02);
		border-radius: var(--size-03);

		width: 2em;
		height: 2em;

		display: flex;
		justify-content: center;
		align-items: center;
		padding: var(--size-02);
		transition: transform 0.1s;
	}

	button:hover {
		background-color: var(--gray-4);
	}

	button:active {
		transform: translateY(1px);
		box-shadow: none;
	}
</style>
