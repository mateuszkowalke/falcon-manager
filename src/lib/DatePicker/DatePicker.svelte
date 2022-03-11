<script lang="ts">
	export let format = 'YYYY-MM-DD';
	export let date: Date;
	export let name: string;

	let internal;

	const input = (x) => {
        internal = `${x.getYear()}-${x.getMonth()}-${x.getDay()}`;
	    console.log(internal);
	};
	const output = (x) => (date = new Date(x));

	$: input(date);
	$: output(internal);
</script>

<input name={name} type="date" bind:value={internal} />
