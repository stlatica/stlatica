<script>
	import { superForm } from 'sveltekit-superforms';
	import { generateUlid } from '$lib/faker/ulid';

	export let form;
	export let data;

	// let { user_id, preferred_user_id, preferred_user_name, is_public, mail_address } = data.props;

	const { form: schema } = superForm(data.form);

	const regenerate_id = () => {
		// schema.set({ ...schema, user_id: ulid() });

		schema.set({ ...$schema, user_id: generateUlid() });
		// user_id = ulid();
	};
</script>

<h2>ユーザー追加</h2>

<form method="post">
	<div class="line">
		<div class="labelname">user_id</div>
		<input class="input" type="text" name="user_id" bind:value={$schema.user_id} />
		<button class="button is-info" type="button" onclick={regenerate_id}>generate id</button>
	</div>

	<label>
		<div class="labelname">preferred_user_id</div>
		<input
			class="input"
			type="text"
			name="preferred_user_id"
			bind:value={$schema.preferred_user_id}
		/>
	</label>

	<label>
		<div class="labelname">preferred_user_name</div>
		<input
			class="input"
			type="text"
			name="preferred_user_name"
			bind:value={$schema.preferred_user_name}
		/>
	</label>

	<label class="checkbox">
		<div class="labelname">is_public</div>
		<input type="checkbox" name="is_public" bind:checked={$schema.is_public} />
	</label>

	<label>
		<div class="labelname">mail_address</div>
		<input class="input" type="email" name="mail_address" bind:value={$schema.mail_address} />
	</label>

	<button class="button is-primary" type="submit">追加</button>
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
	label,
	.line {
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
		background-color: var(--bulma-text-30);
		color: var(--bulma-text-30-invert);
	}

	input[type='text'],
	input[type='email'] {
		flex: 1;
	}

	input[type='checkbox'] {
		margin-left: 0.5em;
	}
</style>
