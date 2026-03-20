<script lang="ts">
	import { page } from '$app/state';
	import type { Snippet } from 'svelte';

	let {
		href,
		target = '_self',
		external = false,
		onclick,
		children
	}: {
		href: string;
		target?: string;
		external?: boolean;
		onclick?: () => void;
		children?: Snippet;
	} = $props();

	let isActive = $derived(!external && page.url.pathname === href);
</script>

<a
	{href}
	{target}
	class="flex flex-row items-center gap-3 w-full h-fit px-2 py-1.5 rounded hover:bg-stroke-light justify-between"
	class:bg-stroke-light={isActive}
	onclick={onclick}
	rel={external ? 'noopener noreferrer' : undefined}
>
	<div class="flex flex-row items-center gap-3">
		{#if children}{@render children()}{/if}
	</div>
	{#if external}
		<img height="9" width="10" src="/images/arrow-angle.svg" alt="External link" loading="lazy" />
	{/if}
</a>
