<script lang="ts">
	import { beforeNavigate } from '$app/navigation';
	import { postPlat } from '$lib/orval/client/stlaticaInternalApi';
	import { fakerJA } from '@faker-js/faker';

	let isFaker = true;

	/**
	 * 実行制御
	 */
	let active = false;

	let user_id = '';

	const intervalID = setInterval(async () => {
		if (!active) {
			return;
		}

		const content = isFaker ? fakerJA.lorem.sentence() : `test plat ${new Date().toISOString()}`;

		const res = await postPlat({
			user_id,
			content
		});

		console.log(res.status);
		console.log(res);
	}, 1000);

	beforeNavigate(() => {
		clearInterval(intervalID);
	});
</script>

<h1>plat定期投稿</h1>

<div class="container">
	<div class="field">
		<label class="label">
			<div>ユーザーID</div>
			<div class="control">
				<input class="input" type="text" name="user_id" bind:value={user_id} />
			</div>
		</label>

		<!-- <div class="field">
			<label class="checkbox">
				<input type="checkbox" bind:checked={isFaker} />
				Faker.jsを使う
			</label>
		</div> -->

		{#if !active}
			<button
				type="button"
				class="button w-full is-primary"
				on:click={() => {
					active = true;
				}}>start</button
			>
		{:else}
			<button
				type="button"
				class="button w-full"
				on:click={() => {
					active = false;
				}}>stop</button
			>
		{/if}
	</div>
</div>

<style>
	.container {
		display: flex;
		flex-direction: column;
		gap: 10px;
	}

	.w-full {
		width: 100%;
	}
</style>
