<script context="module" lang="ts">
	import type { Load } from '@sveltejs/kit';
	export const load: Load = async function ({ fetch }) {
		const url = '/api/falcons';
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
	};
</script>

<script lang="ts">
	import { goto } from '$app/navigation';
	import { Accordion, AccordionItem } from 'svelte-accessible-accordion';
	import type { Falcon, Species, Aviary } from '@prisma/client';

	export let falcons: (Falcon & {
		species: Species;
		aviary: Aviary;
	})[];

	const editFalcon = (id: number) => {
		goto(`/falcons/${id}`);
	};

	const deleteFalcon = (id: number) => {
		console.log(`deleting falcon with id: ${id}`);
	};
</script>

<svelte:head>
	<title>Falcons</title>
</svelte:head>

<h1>Falcons</h1>

<button on:click={() => goto('falcons/new')}>New Falcon</button>

<Accordion>
	{#each falcons as falcon}
		<AccordionItem title={falcon.name}>
			<ul>
				<li>
					<p>Ring: {falcon.ring}</p>
				</li>
				<li>
					<p>{falcon.species.name}</p>
				</li>
				<li>
					<p>{falcon.sex.toLowerCase()}</p>
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
				<li>
					<button on:click={() => editFalcon(falcon.id)}>Edit</button>
				</li>
				<li><button on:click={() => deleteFalcon(falcon.id)}>Delete</button></li>
			</ul>
		</AccordionItem>
	{/each}
</Accordion>
