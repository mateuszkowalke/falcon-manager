<script context="module" lang="ts">
	import type { Load } from '@sveltejs/kit';
	import { loadFalconFormData } from '$lib/loadFalconFormData';
	export const load: Load = async function ({ params, fetch }) {
		const formData = await loadFalconFormData();
		if ('error' in formData) {
			return {
				status: formData.status,
				error: formData.error
			};
		}
		const {species, breedingProjects, aviaries} = formData;
		return {
			props: {
				species,
				breedingProjects,
				aviaries
			}
		};
	};
</script>

<script lang="ts">
	import type { Aviary, Falcon, Species, BreedingProject } from '@prisma/client';
	import FalconForm from '$lib/components/FalconForm/FalconForm.svelte';
	export let falcon: Falcon & {
		species: Species;
		aviary: Aviary;
	} = {} as Falcon & { species: Species; aviary: Aviary };
	export let species: Species[];
	export let breedingProjects: BreedingProject[];
	export let aviaries: Aviary[];
	let error: string;

	async function save() {
		error = '';
		try {
			const url = '/api/falcons';
			const res = await fetch(url, { method: 'POST', body: JSON.stringify(falcon) });
			if (!res.ok) {
				error = 'error saving falcon data';
			}
		} catch (err) {
			console.error(err);
			error = err as string;
		}
	}
</script>

<svelte:head>
	<title>New Falcon</title>
</svelte:head>

<FalconForm {falcon} {species} {breedingProjects} {aviaries} submitHandler={save} />
{#if error}
	<h3>{error}</h3>
{/if}
