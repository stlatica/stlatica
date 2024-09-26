<script lang="ts">
	export let data;
</script>

<main>
	<table>
		<thead>
			<tr>
				<th>user_id</th>
				<!-- <th>plat_id</th> -->
				<th>content</th>
				<th>created_at</th>
			</tr>
		</thead>
		<tbody>
			{#each data.plats as plat}
				<tr>
					<td>{plat.user_id}</td>
					<!-- <td>{plat.plat_id}</td> -->
					<td>{plat.content}</td>
					<td>{plat.created_at}</td>
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
