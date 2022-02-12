<script context="module">
	export async function load({ params, fetch, session, stuff }) {
		const url = 'api/falcons';
		const res = await fetch(url);

		if (res.ok) {
			const falcons = await res.json();
			return {
				props: {
					falcons
				}
			};
		}

		return {
			status: res.status,
			error: new Error(`Could not load ${url}`)
		};
	}
</script>

<script lang="ts">
	export let falcons;
	// todo
    // use context and nested components for accordion
</script>

<svelte:head>
	<title>Falcons</title>
</svelte:head>

<h1>Falcons</h1>
<ul>
	{#each falcons as falcon}
		<li on:click={toggle} class="hidden">
			<h3>{falcon.name}</h3>
			<section>
				<h4>Ring:</h4>
				<p>{falcon.ring}</p>
			</section>
		</li>
	{/each}
</ul>

<style>
	.hidden > section {
		display: none;
	}
</style>
