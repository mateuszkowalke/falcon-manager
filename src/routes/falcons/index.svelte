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
</script>

<svelte:head>
	<title>Falcons</title>
</svelte:head>

<h1>Falcons</h1>
<ul>
	{#each falcons as falcon}
		<li><p>{falcon.name}</p></li>
	{/each}
</ul>

<style>
</style>
