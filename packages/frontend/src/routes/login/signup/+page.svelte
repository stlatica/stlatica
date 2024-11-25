<script lang="ts">
	import { PAGES } from '$lib/routes/ROUTES';
	import Textfield from '$lib/components/form/input/textfield.svelte';
	import * as v from 'valibot';
	import { goto } from '$app/navigation';
	import Button from '$lib/components/material/button/button.svelte';

	const loginSchema = v.object({
		id: v.string(),
		password: v.string()
	});

	let loading = false;
	let message = '';

	const onsubmit = async (e: SubmitEvent) => {};
</script>

<div class="container">
	<h1>Proto Stlatica</h1>
	<h2>Signup</h2>

	<form method="post" {onsubmit}>
		<div class="container">
			<p class="line">
				<label class="textlabel" for="id"> ID (email) </label>
				<Textfield value="hoge" id="id" name="id" />
			</p>

			<p class="line">
				<label class="textlabel" for="password"> Password </label>
				<Textfield value="hoge" id="password" name="password" type="password" />
			</p>

			<!-- <p class="line">
				<label class="textlabel" for="invite"> Invite Code </label>
				<Textfield value="hoge" id="invite" name="invite_code" type="password" />
			</p> -->

			<Button variant="outlined" type="submit">
				{#if loading}
					Loading...
				{:else}
					Register
				{/if}
			</Button>
		</div>
	</form>

	<div>
		<a href={PAGES['/login']}> ログインへ戻る</a>
	</div>

	<div class="red">
		{message}
	</div>
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
