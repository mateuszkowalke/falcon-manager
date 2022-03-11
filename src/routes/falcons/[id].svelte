<script context="module">
	export async function load({ params, fetch, session, stuff }) {
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
	}
</script>

<script lang="ts">
    import type { Falcon } from '@prisma/client'
    import DatePicker from '$lib/DatePicker/DatePicker.svelte'
	export let falcon: Falcon;
	let editing = false;
	let sex;

	function edit() {
		editing = true;
	}

	function save() {
		editing = false;
	}

	function del() {}
</script>

<svelte:head>
	<title>{falcon.name}</title>
</svelte:head>

{#if editing}
	<ul>
		<li>
			<label for="name">Name:</label>
			<input type="text" name="name" bind:value={falcon.name} />
		</li>
		<li>
			<label for="ring">Ring:</label>
			<input type="text" name="ring" bind:value={falcon.ring} />
		</li>
		<li>
			<label for="sex">Sex:</label>
            <select name="sex" bind:value={sex}>
                {#each ['male', 'female', 'unknown'] as opt}
                    <option value={opt}>
                        {opt}
                    </option>
                {/each}
            </select>
		</li>
		<li>
			<label for="birthDate">Birth date:</label>
            <DatePicker bind:date={falcon.birthDate} name="birthDate"/>
		</li>
	</ul>
	<button on:click={save}>Save</button>
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
