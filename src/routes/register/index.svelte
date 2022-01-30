<script lang="ts">
	import { goto } from '$app/navigation';
	import { RolesEnum as Role } from '$lib/roles.enum';

	let name: string;
	let email: string;
	let password: string;
	let password2: string;
	let role: Role;
	let error: string;

	async function register() {
		error = undefined;
		if (password !== password2) {
			error = 'supplied passwords are different';
			return;
		}
		try {
			const res = await fetch('/api/auth/register', {
				method: 'POST',
				body: JSON.stringify({
					name,
					email,
					password,
					role,
					address: {
						street,
						no,
						zipCode,
						city,
						country
					}
				}),
				headers: {
					'Content-Type': 'application/json'
				}
			});
			if (res.ok) {
				//store logged in true - user
				goto('/');
			} else {
				error = 'error registering';
			}
		} catch (err) {
			console.error(err);
			error = err;
		}
	}
</script>

<svelte:head>
	<title>Register</title>
</svelte:head>

<h1>register page</h1>
<label>
	name
	<input type="text" bind:value={name} placeholder="Name" />
</label>
<br />
<label>
	email
	<input type="email" bind:value={email} placeholder="Email" />
</label>
<br />
<label>
	password
	<input type="password" bind:value={password} placeholder="Password" />
</label>
<br />
<label>
	repeat password
	<input type="password" bind:value={password2} placeholder="Repeat Password" />
</label>
<br />
<label>
	<input type="radio" bind:group={role} value={Role.BREEDER} />
	Breeder
</label>
<br />
<label>
	<input type="radio" bind:group={role} value={Role.WORKER} />
	Worker
</label>
<br />
<button on:click={register}>register</button>
{#if error}
	<h3>{error}</h3>
{/if}

<style>
</style>
