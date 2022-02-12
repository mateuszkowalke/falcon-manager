<script lang="ts">
	import { goto } from '$app/navigation';
	import { session } from '$app/stores';

    async function logout() {
        await fetch('/api/auth/logout', {
            method: 'POST'
        });
        $session.user = null;
        goto('/');
    }
    logout();
</script>
