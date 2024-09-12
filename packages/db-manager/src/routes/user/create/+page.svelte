<script>
	import Button from '$lib/material/button/button.svelte';
	import MyInput from '$lib/material/input/text-input.svelte';
	import { ulid } from 'ulidx';
	import { superForm } from 'sveltekit-superforms';

	export let form;
	export let data;

	// let { user_id, preferred_user_id, preferred_user_name, is_public, mail_address } = data.props;

	const { form: schema } = superForm(data.form);

	const regenerate_id = () => {
		// schema.set({ ...schema, user_id: ulid() });

		schema.set({ ...$schema, user_id: ulid() });
		// user_id = ulid();
	};

	$: console.log($schema);
</script>

<h2>ユーザー追加</h2>

<form method="post">
	<div class="line">
		<div class="field border label" style="flex: 1">
			<input type="text" name="user_id" bind:value={$schema.user_id} />
			<label for="">UserID</label>
		</div>
		<button class="small-round" type="button" onclick={regenerate_id}>generate id</button>
	</div>

	<div class="line">
		<div class="field border label" style="flex: 1">
			<input type="text" name="preferred_user_id" bind:value={$schema.preferred_user_id} />
			<label for="">Preferred UserID</label>
		</div>
	</div>

	<div class="line">
		<div class="field border label" style="flex: 1">
			<input type="text" name="preferred_user_name" bind:value={$schema.preferred_user_name} />
			<label for="">Preferred UserName</label>
		</div>
	</div>

	<div class="line">
		<div class="field border label" style="flex: 1">
			<input type="email" name="mail_address" bind:value={$schema.mail_address} />
			<label for="">Mail Address</label>
		</div>
	</div>

	<div>
		<label class="checkbox">
			<input type="checkbox" name="is_public" bind:checked={$schema.is_public} />
			<span>is_public</span>
		</label>
	</div>

	<button class="outline" type="submit">保存</button>

	<!-- <div class="label_">
		<div class="labelname">user_id</div>
		<MyInput type="text" name="user_id" bind:value={$schema.user_id} />
		<button class="outline" type="button" onclick={regenerate_id}>generate id</button>
	</div>

	<label>
		<div class="labelname">preferred_user_id</div>
		<input type="text" name="preferred_user_id" bind:value={$schema.preferred_user_id} />
	</label>

	<label>
		<div class="labelname">preferred_user_name</div>
		<MyInput type="text" name="preferred_user_name" bind:value={$schema.preferred_user_name} />
	</label>

	<label>
		<div class="labelname">is_public</div>
		<input type="checkbox" name="is_public" bind:checked={$schema.is_public} />
	</label>

	<label>
		<div class="labelname">mail_address</div>
		<MyInput type="email" name="mail_address" bind:value={$schema.mail_address} />
	</label>

	<button class="outline" type="submit">保存</button> -->
</form>

{#if form}
	<div class="result">
		<p>api result</p>
		<!-- <p>success: {form.success}</p>
		<p>data:</p>
		<pre>{JSON.stringify(form?.data, null, 2)}</pre> -->

		<pre>{form.message}</pre>

		<!-- <pre>{JSON.stringify(form, null, 2)}</pre> -->
	</div>
{/if}

<style>
	h2 {
		text-align: center;
		padding-bottom: 2rem;
	}

	.line {
		display: flex;
		/* background-color: red; */
	}

	.result {
		margin-top: var(--size-15);
		padding: var(--font-size-sm);
		background-color: var(--gray-7);
	}
</style>
