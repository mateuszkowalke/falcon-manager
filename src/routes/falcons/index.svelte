<script context="module">
    export async function load({ params, fetch, session, stuff }) {
		const url = 'api/falcons';
		const res = await fetch(url);

		if (res.ok) {
			return {
				props: {
					falcons: await res.json()
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
    let falcons;
</script>

<svelte:head>
	<title>Falcons</title>
</svelte:head>

<h1>Falcons</h1>
<p>{falcons}</p>

<style>
</style>

