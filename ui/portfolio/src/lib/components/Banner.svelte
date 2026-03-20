<script lang="ts">
	import type { Snippet } from 'svelte';

	let { border = false, children }: { border?: boolean; children?: Snippet } = $props();
</script>

<div class="w-full min-h-[50px]" class:border-b={border} class:border-b-stroke={border}>
	{#if children}{@render children()}{/if}
</div>
