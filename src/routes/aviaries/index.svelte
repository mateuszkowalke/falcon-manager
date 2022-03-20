<script context="module" lang="ts">
	import type { Load } from '@sveltejs/kit';
	export const load: Load = async function ({ fetch }) {
		const url = '/api/aviaries';
		const res = await fetch(url);

		if (res.ok) {
			const aviaries = await res.json();
			return {
				props: {
					aviaries
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
	import { goto } from '$app/navigation';
	import { Accordion, AccordionItem } from 'svelte-accessible-accordion';
	import type { Aviary } from '@prisma/client';

	export let aviaries: Aviary[];

	const editAviary = (id: number) => {
		goto(`/aviaries/${id}`);
	};

	const deleteAviary = (id: number) => {
		console.log(`deleting aviary with id: ${id}`);
	};
</script>

<svelte:head>
	<title>Aviaries</title>
</svelte:head>

<h1>Aviaries</h1>

<Accordion>
	{#each aviaries as aviary}
		<AccordionItem title={aviary.name}>
			<ul>
				<li>
					<p>Capacity: {aviary.capacity}</p>
				</li>
				<li>
					<p>Last cleaned: {aviary.lastCleaned}</p>
				</li>
				<li>
					<button on:click={() => editAviary(aviary.id)}>Edit</button>
				</li>
				<li><button on:click={() => deleteAviary(aviary.id)}>Delete</button></li>
			</ul>
		</AccordionItem>
	{/each}
</Accordion>
