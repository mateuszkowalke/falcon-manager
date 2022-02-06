<script lang="ts">
	import Header from '$lib/header/Header.svelte';
	import Footer from '$lib/footer/Footer.svelte';
	import '../app.css';
</script>

<svelte:head>
	<title>Falcon Manager</title>
</svelte:head>

<Header />

<main>
	<slot />
</main>

<Footer />
