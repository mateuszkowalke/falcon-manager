<script context="module" lang="ts">
	import type { Load } from '@sveltejs/kit';
	export const load: Load = async function ({ params, fetch }) {
		const url = `/api/aviaries/${params.id}`;
		const res = await fetch(url);

		if (res.ok) {
			const aviary = await res.json();
			return {
				props: {
					aviary
				}
			};
		}

		return {
			status: res.status,
			error: new Error(`Could not load ${url}`)
		};
	};
</script>

<script lang="ts">
	import type { Falcon, Aviary } from '@prisma/client';
	import DatePicker from '$lib/DatePicker/DatePicker.svelte';
	export let aviary: Aviary & {
		falcons: Falcon[];
	};
	let editing = false;

	function edit() {
		editing = true;
	}

	function save() {
		editing = false;
	}

	function del() {}
</script>

<svelte:head>
	<title>{aviary.name}</title>
</svelte:head>

{#if editing}
	<form>
		<label for="name">Name:</label>
		<input type="text" name="name" bind:value={aviary.name} />
		<label for="lastCleaned">Last cleaned:</label>
		<DatePicker name="lastCleaned" bind:date={aviary.lastCleaned} />
		<button type="submit" on:click|preventDefault={save}>Save</button>
	</form>
{:else}
	<h1>Aviary {aviary.name}</h1>

	<ul>
		<li>
			<p>Capacity: {aviary.capacity}</p>
		</li>
		<li>
			<p>Last cleaned: {aviary.lastCleaned}</p>
		</li>
		<li>
			Falcons in aviary:
			<ul>
				{#each aviary.falcons as falcon}
					<li>{falcon.name}</li>
				{/each}
			</ul>
		</li>
	</ul>
	<button on:click={edit}>Edit</button>
	<button>Delete</button>
{/if}
