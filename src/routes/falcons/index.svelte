<script context="module">
	export async function load({ params, fetch, session, stuff }) {
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
	}
</script>

<script lang="ts">
	import { Accordion, AccordionItem } from 'svelte-accessible-accordion';

	export let falcons;
</script>

<svelte:head>
	<title>Falcons</title>
</svelte:head>

<h1>Falcons</h1>

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
                    <p>In aviary: {falcon.aviary?.name || "not specified"}</p>
                </li>
                <li>
                    <a href={`/falcons/${falcon.id}`}>Details</a>
                </li>
            </ul>
        </AccordionItem>
    {/each}
</Accordion>
