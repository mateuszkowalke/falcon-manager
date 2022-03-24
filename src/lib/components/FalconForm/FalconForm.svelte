<script lang="ts">
	import DatePicker from '$lib/DatePicker/DatePicker.svelte';
	import type { Sex, Aviary, Falcon, Species, BreedingProject } from '@prisma/client';

	export let falcon: Falcon & {
		species: Species;
		aviary: Aviary;
	};
	export let species: Species[];
	export let breedingProjects: BreedingProject[];
	export let aviaries: Aviary[];
	export let submitHandler: () => void;
</script>

<form>
	<label for="name">Name:</label>
	<input type="text" name="name" bind:value={falcon.name} />
	<label for="ring">Ring:</label>
	<input type="text" name="ring" bind:value={falcon.ring} />
	<label for="sex">Sex:</label>
	<select name="sex" bind:value={falcon.sex}>
		{#each ['male', 'female', 'unknown'] as opt}
			<option value={opt.toUpperCase()}>
				{opt}
			</option>
		{/each}
	</select>
	<label for="species">Species:</label>
	<select name="species" bind:value={falcon.speciesId}>
		{#each species as spec}
			<option value={spec.id}>
				{spec.name}
			</option>
		{/each}
	</select>
	<label for="breedingProjects">BreedingProject:</label>
	<select name="breedingProjects" bind:value={falcon.breedingProjectId}>
		{#each breedingProjects as breedingProject}
			<option value={breedingProject.id}>
				{breedingProject.name}
			</option>
		{/each}
	</select>
	<label for="aviaries">Aviary:</label>
	<select name="aviaries" bind:value={falcon.aviaryId}>
		{#each aviaries as aviary}
			<option value={aviary.id}>
				{aviary.name}
			</option>
		{/each}
	</select>
	<label for="birthDate">Birth date:</label>
	<DatePicker name="birthDate" bind:date={falcon.birthDate} />
	<label for="source">Source:</label>
	<input type="text" name="source" bind:value={falcon.source} />
	<label for="notes">Notes:</label>
	<input type="text" name="notes" bind:value={falcon.notes} />
	<button type="submit" on:click|preventDefault={submitHandler}>Save</button>
</form>
