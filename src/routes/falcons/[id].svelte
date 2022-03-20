<script context="module" lang="ts">
	import type { Load } from '@sveltejs/kit';
	export const load: Load = async function ({ params, fetch }) {
		const url = `/api/falcons/${params.id}`;
		const res = await fetch(url);

		if (res.ok) {
			const falcon = await res.json();
			return {
				props: {
					falcon
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
	import type { Falcon, Sex, Species, Aviary } from '@prisma/client';
	import DatePicker from '$lib/DatePicker/DatePicker.svelte';
	export let falcon: Falcon & {
		species: Species;
		aviary: Aviary;
	};
	let editing = false;
	let sex: Sex;
	let error: string;

	function edit() {
		editing = true;
	}

	async function save() {
		error = undefined;
		try {
			const url = `/api/falcons/${falcon.id}`;
			const res = await fetch(url, { method: 'PATCH', body: JSON.stringify(falcon) });
			if (res.ok) {
				editing = false;
			} else {
				error = 'error updating falcon data';
			}
		} catch (err) {
			console.error(err);
			error = err;
		}
	}

	function del() {}
</script>

<svelte:head>
	<title>{falcon.name}</title>
</svelte:head>

{#if editing}
	<form>
		<label for="name">Name:</label>
		<input type="text" name="name" bind:value={falcon.name} />
		<label for="ring">Ring:</label>
		<input type="text" name="ring" bind:value={falcon.ring} />
		<label for="sex">Sex:</label>
		<select name="sex" bind:value={sex}>
			{#each ['male', 'female', 'unknown'] as opt}
				<option value={opt}>
					{opt}
				</option>
			{/each}
		</select>
		<label for="birthDate">Birth date:</label>
		<DatePicker name="birthDate" bind:date={falcon.birthDate} />
		<label for="source">Source:</label>
		<input type="text" name="source" bind:value={falcon.source} />
		<label for="notes">Notes:</label>
		<input type="text" name="notes" bind:value={falcon.notes} />
		<button type="submit" on:click|preventDefault={save}>Save</button>
	</form>
{:else}
	<h1>Falcon {falcon.name}</h1>

	<ul>
		<li>
			<p>Ring: {falcon.ring}</p>
		</li>
		<li>
			<p>Species: {falcon.species.name}</p>
		</li>
		<li>
			<p>Sex: {falcon.sex.toLowerCase()}</p>
		</li>
		<li>
			<p>Hatched: {falcon.birthDate}</p>
		</li>
		<li>
			<p>Source: {falcon.source}</p>
		</li>
		<li>
			<p>In aviary: {falcon.aviary?.name || 'not specified'}</p>
		</li>
	</ul>
	<button on:click={edit}>Edit</button>
	<button>Delete</button>
{/if}
