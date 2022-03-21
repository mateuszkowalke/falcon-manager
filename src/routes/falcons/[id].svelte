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
	import type { Falcon, Species, Aviary } from '@prisma/client';
	import FalconForm from '$lib/components/FalconForm/FalconForm.svelte';
	export let falcon: Falcon & {
		species: Species;
		aviary: Aviary;
	};
	let error: string;

	async function update() {
		error = undefined;
		try {
			const url = `/api/falcons/${falcon.id}`;
			const res = await fetch(url, { method: 'PATCH', body: JSON.stringify(falcon) });
			if (!res.ok) {
				error = 'error updating falcon data';
			}
		} catch (err) {
			console.error(err);
			error = err;
		}
	}
</script>

<svelte:head>
	<title>Falcon details</title>
</svelte:head>

<h1>Falcon {falcon.name}</h1>
<FalconForm {falcon} submitHandler={update} />
{#if error}
	<h3>{error}</h3>
{/if}
