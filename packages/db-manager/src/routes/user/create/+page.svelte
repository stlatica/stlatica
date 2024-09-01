<script>
	import Button from '$lib/material/button/button.svelte';
	import MyInput from '$lib/material/input/text-input.svelte';
	import { ulid } from 'ulidx';

	// let user_id = ulid();
	// let preferred_user_id = fakerJA.string.alphanumeric({ length: { min: 8, max: 16 } });
	// let preferred_user_name = fakerJA.person.fullName();
	// let is_public = true;
	// let mail_address = 'test@example.com';
	export let form;
	export let data;

	let { user_id, preferred_user_id, preferred_user_name, is_public, mail_address } = data.props;

	const regenerate_id = () => {
		user_id = ulid();
	};
</script>

<h2>ユーザー追加</h2>

<form method="post">
	<div class="label">
		<div class="labelname">user_id</div>
		<MyInput type="text" name="user_id" readonly bind:value={user_id} />
		<Button type="button" onclick={regenerate_id}>generate id</Button>
	</div>

	<label>
		<div class="labelname">preferred_user_name</div>
		<MyInput type="text" name="preferred_user_id" bind:value={preferred_user_id} />
	</label>

	<label>
		<div class="labelname">preferred_user_id</div>
		<MyInput type="text" name="preferred_user_name" bind:value={preferred_user_name} />
	</label>

	<label>
		<div class="labelname">is_public</div>
		<input type="checkbox" name="is_public" bind:checked={is_public} />
	</label>

	<label>
		<div class="labelname">mail_address</div>
		<MyInput type="email" name="mail_address" bind:value={mail_address} />
	</label>

	<Button type="submit">保存</Button>
</form>

<div class="result">
	<p>api result</p>
	{#if form}
		<p>success: {form.success}</p>
		<p>data:</p>
		<pre>{JSON.stringify(form?.data, null, 2)}</pre>

		<p>message: {form.message}</p>
	{/if}

	<!-- <pre>{JSON.stringify(form, null, 2)}</pre> -->
</div>

<style>
	h2 {
		text-align: center;
		padding-bottom: 2rem;
	}

	label,
	.label {
		display: flex;
		align-items: center;
		gap: var(--font-size-lg);
	}

	.labelname {
		width: 8rem;
		font-size: var(--font-size-sm);
	}

	.result {
		margin-top: var(--size-15);
		padding: var(--font-size-sm);
		background-color: var(--gray-7);
	}
</style>
