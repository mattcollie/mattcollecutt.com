<script lang="ts">
	import type { Snippet } from 'svelte';

	let { onclick, children, class: className, ...rest }: { onclick?: () => void; children?: Snippet; class?: string; [key: string]: unknown } = $props();
</script>

<button onclick={onclick} class={className} {...rest}>
	{#if children}{@render children()}{/if}
</button>
