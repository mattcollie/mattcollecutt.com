<script lang="ts">
	import type { Snippet } from 'svelte';

	let { title = '', children }: { title?: string; children?: Snippet } = $props();
</script>

<div class="flex flex-col p-4 gap-1">
	{#if title}<span class="text-xs font-semibold">{title}</span>{/if}
	{#if children}{@render children()}{/if}
</div>
