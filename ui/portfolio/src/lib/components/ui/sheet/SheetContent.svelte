<script lang="ts">
	import { cn } from '$lib/utils';
	import type { Snippet } from 'svelte';

	let {
		open = false,
		side = 'left',
		onclose,
		children,
		class: className
	}: {
		open?: boolean;
		side?: 'left' | 'right';
		onclose?: () => void;
		children?: Snippet;
		class?: string;
	} = $props();
</script>

{#if open}
	<!-- Overlay -->
	<button
		class="fixed inset-0 z-40 bg-black/60"
		onclick={onclose}
		aria-label="Close navigation"
	></button>

	<!-- Panel -->
	<div
		class={cn(
			'fixed inset-y-0 z-50 flex flex-col bg-secondary border-r border-r-stroke',
			side === 'left' ? 'left-0' : 'right-0',
			'w-3/4 sm:w-1/2 md:w-56',
			'animate-slide-in',
			className
		)}
	>
		{#if children}{@render children()}{/if}
	</div>
{/if}

<style>
	@keyframes slide-in {
		from {
			transform: translateX(-100%);
		}
		to {
			transform: translateX(0);
		}
	}

	:global(.animate-slide-in) {
		animation: slide-in 0.3s ease-in-out;
	}
</style>
