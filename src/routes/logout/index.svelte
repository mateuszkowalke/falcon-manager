<script lang="ts">
	import { goto } from '$app/navigation';
	import { session } from '$app/stores';

	fetch('/api/auth/logout', {
		method: 'POST'
	});
	$session.user = null;
	goto('/');
</script>
