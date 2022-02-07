<script context="module">
	export async function load({ session }) {
		if (session.user) {
			return {
				status: 302,
				redirect: '/'
			};
		}
		return {};
	}
</script>

<script lang="ts">
	import { goto } from '$app/navigation';
	import { session } from '$app/stores';

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
					password
				}),
				headers: {
					'Content-Type': 'application/json'
				}
			});
			if (res.ok) {
				const body = await res.json();
				$session.user = {
					name: body.user.name,
					email: body.user.email,
					image: body.user.photoFile
				};
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
