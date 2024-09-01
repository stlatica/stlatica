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

	table,
	table th,
	table td {
		border: 1px solid var(--gray-10);
	}
</style>
