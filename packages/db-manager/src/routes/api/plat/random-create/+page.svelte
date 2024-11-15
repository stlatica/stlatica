<script lang="ts">
	import { AsyncQueue } from '$lib/AsyncQueue';
	import { postPlat } from '$lib/orval/client/stlaticaInternalApi';
	import { fakerJA } from '@faker-js/faker';

	export let data;
	let isFaker = true;

	let progress: string = '待機中';

	async function PostMany(num: number) {
		progress = '初期化中';

		const array = new Array(num).fill(0).map((_, i) => {
			const content = isFaker ? fakerJA.lorem.sentence() : `test plat ${i + 1}`;

			return {
				user_id: data.users[Math.floor(Math.random() * data.users.length)].user_id,
				content
			};
		});

		let cnt = 0;
		await AsyncQueue(
			array,
			async (x, i) => {
				await postPlat(x);
				progress = `${cnt++} / ${num}`;
			},
			1000
		);

		progress = '完了';
	}
</script>

<h1>plat大量投稿</h1>

<div class="container">
	<button class="button" on:click={() => PostMany(10)}>10回実行</button>
	<button class="button" on:click={() => PostMany(100)}>100回実行</button>
	<button class="button" on:click={() => PostMany(1000)}>1000回実行</button>
	<button class="button" on:click={() => PostMany(10000)}>10000回実行</button>
	<button class="button" on:click={() => PostMany(100000)}>100000回実行</button>
	<label>
		Faker.jsを使う
		<input type="checkbox" bind:checked={isFaker} />
	</label>

	progress: {progress}
</div>

<style>
	.container {
		display: flex;
		flex-direction: column;
		gap: 10px;
	}
</style>
