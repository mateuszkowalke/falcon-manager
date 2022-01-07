<script lang="ts">
	export let user: any;
    console.log('user in nav: ' + user)
</script>

<header>
	<nav>
		<ul>
			{#if user}
				<li><p>{user.name}</p></li>
			{:else}
				<li><a href="login">login</a></li>
				<li><a href="register">register</a></li>
			{/if}
		</ul>
	</nav>
</header>

<style>
</style>
