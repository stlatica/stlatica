<script lang="ts">
	type Props = {
		username: string;
		userid: string;
		content: string;
	};

	let { username, userid, content }: Props = $props();

	import UserIcon from "./user-icon.svelte";
</script>

<div class="container">
	<div class="info">
		<div class="icon">
			<UserIcon src="https://via.placeholder.com/100/b7794a/fff.webp"></UserIcon>
		</div>

		<div class="ut-flex-v">
			<div class="username">{username}</div>
			<div class="userid">@{userid}</div>
		</div>
	</div>

	<div class="content">{content}</div>
</div>

<style>
	.container {
		width: 100%;
		height: 100%;
		padding: var(--size-05);
		border: solid 1px var(--gray-8);
	}

	.username {
		font-size: var(--font-size-5);
		font-weight: bold;
	}

	.userid {
		color: var(--gray-11);
		font-size: var(--font-size-2);
	}

	.info {
		display: flex;
		flex-direction: columns;
		gap: var(--size-04);
	}

	.content {
		margin: var(--font-size-2);
		/* アイコンと同幅 */
		margin-left: 60px;
	}

	.icon {
		width: 60px;
	}
</style>
