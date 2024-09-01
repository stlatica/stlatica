<script lang="ts">
	export let data;
</script>

<main>
	<table>
		<thead>
			<tr>
				<th>user_id</th>
				<th>preferred_user_id</th>
				<th>preferred_user_name</th>
				<th>is_public</th>
				<th>mail_address</th>
			</tr>
		</thead>
		<tbody>
			{#each data.users as user}
				<tr>
					<td>{user.user_id}</td>
					<td>{user.preferred_user_id}</td>
					<td>{user.preferred_user_name}</td>
					<td>{user.is_public}</td>
					<td>{user.mail_address}</td>
				</tr>
			{/each}
		</tbody>
	</table>
</main>

<style>
	thead {
		font-size: var(--font-size-sm);
	}

	table {
		text-align: center;
		border-collapse: collapse;
	}

	table th,
	table td {
		border-bottom: 2px solid var(--gray-10);
		padding: var(--size-3);
		padding-left: var(--size-5);
	}
</style>
