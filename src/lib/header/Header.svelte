<script lang="ts">
    import { session } from '$app/stores'
</script>

<header>
	<nav>
		<ul>
			{#if $session.user}
				<li><p>{$session.user.name}</p></li>
				<li><a href="/falcons">falcons</a></li>
				<li><a href="/aviaries">aviaries</a></li>
				<li><a href="/logout">logout</a></li>
			{:else}
				<li><a href="/login">login</a></li>
				<li><a href="/register">register</a></li>
			{/if}
		</ul>
	</nav>
</header>

<style>
</style>
