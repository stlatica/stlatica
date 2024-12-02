<script lang="ts">
	import FollowButton from "../button/follow-button.svelte";
	import UserIcon from "./user-icon.svelte";

	let userID = "user_id";
	let isFollow = false;

	const handleFollow = () => {
		isFollow = !isFollow;
	};

	import LetsIconsMeatballsMenu from "~icons/lets-icons/meatballs-menu";
</script>

<div class={"container"}>
	<img class={"header"} alt="header" src="https://via.placeholder.com/700/79b74a/fff.webp" />

	<div class={"parent"}>
		<div class={"icon"}>
			<UserIcon src="https://via.placeholder.com/100/b7794a/fff.webp" />
		</div>
	</div>

	<div class={"toolPalette"}>
		<FollowButton {isFollow} onclick={handleFollow}></FollowButton>
		<LetsIconsMeatballsMenu></LetsIconsMeatballsMenu>
	</div>

	<div class={"texts"}>
		<div class="username">User Name Hereeeeeeeee</div>
		<div>@{userID}</div>
		<div class={"FollowsLine"}>
			<div>100 follower</div>
			<div>100 follow</div>
			<div>100 post</div>
		</div>
		<div>
			ここに自己紹介をかきまああああああああああああああああああああああああああああああああああああす文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数
		</div>
	</div>
</div>

<style>
	.container {
		width: 100%;
		height: 100%;
		/* backgroundColor: var(--mantine-color-gray-9); */
	}

	.username {
		font-size: var(--font-size-6);
		font-weight: bold;
	}

	.header {
		height: 200px;
		width: 100%;
		object-fit: cover;
		display: flex;
		justify-content: center;
		background-color: black;
	}

	.parent {
		position: relative;
		margin-bottom: 0.5rem;
		margin-left: 20px; /* Assuming width is always 100, so 100 / 5 = 20px */
	}

	.icon {
		width: 100px;
		padding: 2px;
		/* margin: 3px; */
		background-color: var(--gray-2);
		border-radius: 50%;
		position: absolute;
		top: -50px; /* Assuming width is always 100, so 100 / 2 = 50px */
	}

	.toolPalette {
		height: 50px; /* Assuming width is always 100, so 100 / 2 = 50px */
		display: flex;
		flex-direction: row;
		gap: 1rem;
		justify-content: flex-end;
		margin-right: 1rem;
		align-items: center;
	}

	.texts {
		margin: 1rem;
		display: flex;
		flex-direction: column;
	}

	.FollowsLine {
		display: flex;
		flex-direction: row;
		align-items: center;
		gap: 1rem;
		margin-top: 0.5rem;
		margin-bottom: 0.5rem;
	}
</style>
