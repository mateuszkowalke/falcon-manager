<script context="module" lang="ts">
	export async function load({ page, fetch, session, stuff }) {
		console.log('in load func');
		if (!session) return {};
		console.log('nav load session: ' + session);
		return {
			props: {
				user: session.user
			}
		};
	}
</script>

<script lang="ts">
	import Header from '$lib/header/Header.svelte';
	import Footer from '$lib/footer/Footer.svelte';
	import '../app.css';

	export let user;
</script>

<svelte:head>
	<title>Falcon Manager</title>
</svelte:head>

<Header {user} />

<main>
	<slot {user} />
</main>

<Footer />
