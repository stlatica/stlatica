<script lang="ts">
	import { superForm } from 'sveltekit-superforms';
	export let data;

	const { form: schema } = superForm(data.form);
</script>

<h2>plat投稿</h2>

<form method="post">
	<div class="field">
		<label class="label">
			<div>ユーザーID</div>
			<div class="control">
				<input class="input" type="text" name="user_id" bind:value={$schema.user_id} />
			</div>
		</label>
	</div>
	<div class="field">
		<label class="label">
			<div>コンテンツ</div>
			<div class="control">
				<textarea class="textarea" name="content" rows="5" bind:value={$schema.content}></textarea>
			</div>
		</label>
	</div>
	<button class="button is-primary" type="submit">送信</button>
</form>

<pre>
{JSON.stringify(data.pickupUser, null, 2)}
</pre>
