<script lang="ts">
	import DatePicker from '$lib/DatePicker/DatePicker.svelte';
	import type { Aviary, Falcon, Species } from '@prisma/client';
	import type { Sex } from '@prisma/client';

	export let falcon: Falcon & {
		species: Species;
		aviary: Aviary;
	};
	export let submitHandler: () => void;
	let sex: Sex;
</script>

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
	<button type="submit" on:click|preventDefault={submitHandler}>Save</button>
</form>
