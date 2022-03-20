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
	import { Accordion, AccordionItem } from 'svelte-accessible-accordion';
	import type { Aviary } from '@prisma/client';

    export let aviaries: Aviary[];
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
					<a href={`/aviaries/${aviary.id}`}>Details</a>
				</li>
			</ul>
		</AccordionItem>
	{/each}
</Accordion>
