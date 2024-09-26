<script>
	import '$lib/styles/tw-variables/_index.css';
	import '$lib/styles/theme.css';
	import '$lib/styles/global.css';
	import '$lib/styles/utilities.css';
</script>

<main>
	<div class="menu">
		<div>DB直操作</div>
		<a href="/user/create"> ユーザー追加 </a>
		<hr />

		<div>DB Viewer</div>
		<a href="/view/plats"> Plat 一覧 </a>
		<a href="/view/users"> ユーザー一覧 </a>
		<hr />

		<div>API経由操作</div>
		<!-- <div>TODO...</div> -->
		<a href="/api/plat/random-create">api</a>
	</div>

	<div class="container">
		<div class="content">
			<slot></slot>
		</div>
	</div>
</main>

<style>
	main {
		display: flex;
		height: 100%;
		/* min-width: 800px; */
	}

	a {
		display: block;
	}

	.menu {
		background-color: var(--gray-3);
		/* height: 100%; */
		padding: 0.5rem;
		padding-right: 0.5rem;
		width: 10rem;
		/* display: flex;
		flex-direction: column; */
	}

	.container {
		flex: auto;
		display: flex;
		flex-direction: column;
		align-items: center;
		overflow: auto;
	}

	.content {
		min-width: 600px;
		/* max-width: 1440px; */
		padding: 1rem;
		display: flex;
		flex-direction: column;
		/* min-width: 15rem; */
		/* background-color: red; */
	}

	hr {
		border-color: var(--gray-8);
	}
</style>
