<script context="module">
	export async function load({ params, fetch, session, stuff }) {
		console.log(params);
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
	export let falcon;
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
			<p>Name:</p>
			<input type="text" bind:value={falcon.name} />
		</li>
		<li>
			<p>Ring:</p>
			<input type="text" bind:value={falcon.ring} />
		</li>
		<li>
			<p>Sex:</p>
            <select bind:value={sex}>
                {#each ['male', 'female', 'unknown'] as opt}
                    <option value={opt}>
                        {opt}
                    </option>
                {/each}
            </select>
		</li>
		<li>
			<p>Hatched:</p>
			<input type="datetime-local" bind:value={falcon.birthDate} />
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
	</ul>
	<button on:click={edit}>Edit</button>
	<button>Delete</button>
{/if}
