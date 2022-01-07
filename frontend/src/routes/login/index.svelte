<script lang="ts">
	import { goto } from '$app/navigation';

	let name: string;
	let password: string;
	let error: string;

	async function login() {
		error = undefined;
		try {
			const res = await fetch('/api/auth/login', {
				method: 'POST',
				body: JSON.stringify({
					name,
					password,
				}),
				headers: {
					'Content-Type': 'application/json'
				}
			});
			if (res.ok) {
				//store logged in true - user
				goto('/');
			} else {
				error = 'error logging in';
			}
		} catch (err) {
			console.error(err);
			error = err;
		}
	}
</script>

<svelte:head>
	<title>Login</title>
</svelte:head>

<h1>login page</h1>
<label>
	name
	<input type="text" bind:value={name} placeholder="Name" />
</label>
<br />
<label>
	password
	<input type="password" bind:value={password} placeholder="Password" />
</label>
<br />
<button on:click={login}>login</button>
{#if error}
	<h3>{error}</h3>
{/if}

<style>
</style>
