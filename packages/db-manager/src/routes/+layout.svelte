<script>
	import '$lib/styles/tw-variables/_index.css';
	import '$lib/styles/global.css';
	import '$lib/styles/utilities.css';
	import 'bulma/css/bulma.min.css';
</script>

<main>
	<div class="menu">
		<p>DB直操作</p>
		<a href="/user/create"> ユーザー追加 </a>
		<hr />

		<p>DB Viewer</p>
		<a href="/view/plats"> Plat 一覧 </a>
		<a href="/view/users"> ユーザー一覧 </a>
		<hr />

		<p>API経由操作</p>
		<!-- <div>TODO...</div> -->
		<a href="/api/plat/random-create">ランダムplat作成</a>
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
		/* background-color: var(--bulma-text-20-invert); */
		color: var(--bulma-text-bold);
		background-color: var(--bulma-text-bold-invert);
		/* height: 100%; */
		padding: 10px;
		padding-right: 20px;
		width: 150px;
		word-break: keep-all;
		/* display: flex;
		flex-direction: column; */
	}

	.container {
		flex: auto;
		/* display: flex; */
		/* flex-direction: column; */
		/* align-items: center; */
		overflow: auto;
	}

	.content {
		min-width: 440px;
		/* max-width: 1440px; */
		padding: 1rem;
		/* display: flex; */
		/* flex-direction: column; */
		/* min-width: 15rem; */
		/* background-color: red; */
	}

	hr {
		--hrcolor: var(--bulma-text-base);
		border-color: var(--hrcolor);
		background-color: var(--hrcolor);
	}
</style>
