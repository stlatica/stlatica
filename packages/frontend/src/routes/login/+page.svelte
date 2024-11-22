<script lang="ts">
	import Textfield from '$lib/components/form/input/textfield.svelte';
	import * as v from 'valibot';
	// import { goto } from '$app/navigation';
	import { PAGES } from '$lib/routes/ROUTES';
	import Button from '$lib/components/material/button/button.svelte';

	let loading = false;
	let message = '';

	const loginSchema = v.object({
		id: v.string(),
		password: v.string()
	});

	const onsubmit = async (e: SubmitEvent) => {};
</script>

<div class="container">
	<h1>Proto Stlatica</h1>
	<h2>Login</h2>

	<form method="post" {onsubmit}>
		<div class="container">
			<p class="line">
				<label class="textlabel" for="id"> ID </label>
				<!-- <input id="id" name="id" /> -->
				<Textfield id="id" name="id"></Textfield>
			</p>
			<p class="line">
				<label class="textlabel" for="password"> Password </label>
				<Textfield id="password" name="password" type="password" />
			</p>

			<div class="red">
				{message}
			</div>

			<Button type="submit" fullwidth>
				{#if loading}
					Loading...
				{:else}
					Login
				{/if}
			</Button>
		</div>
	</form>

	<div>
		<a href={PAGES['/login/signup']}>登録はこちら</a>
	</div>
	<div>パスワードの再設定はこちら(TODO)</div>
</div>

<style>
	h1 {
		text-align: center;
	}

	.red {
		color: var(--st-color-text-red);
	}

	label {
		display: block;
	}

	.line {
		display: flex;
		flex-direction: column;
	}

	.container {
		display: flex;
		flex-direction: column;
		gap: 1em;
		width: 100%;
		align-items: center;
	}
</style>
