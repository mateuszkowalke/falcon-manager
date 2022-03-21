<script lang="ts">
	import type { Aviary, Falcon, Species } from '@prisma/client';
	import FalconForm from '$lib/components/FalconForm/FalconForm.svelte';
	export let falcon: Falcon & {
		species: Species;
		aviary: Aviary;
	} = {} as Falcon & { species: Species; aviary: Aviary };
	let error: string;

	async function save() {
		error = undefined;
		try {
			const url = '/api/falcons';
			const res = await fetch(url, { method: 'POST', body: JSON.stringify(falcon) });
			if (!res.ok) {
				error = 'error saving falcon data';
			}
		} catch (err) {
			console.error(err);
			error = err;
		}
	}
</script>

<svelte:head>
	<title>New Falcon</title>
</svelte:head>

<FalconForm {falcon} submitHandler={save} />
{#if error}
	<h3>{error}</h3>
{/if}
