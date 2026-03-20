<script lang="ts">
	import { cn } from '$lib/utils';
	import type { ButtonProps } from './index';

	let { variant = 'default', size = 'default', class: className, children, ...rest }: ButtonProps = $props();

	const variants: Record<string, string> = {
		default: 'bg-foreground text-background hover:bg-foreground/90',
		ghost: 'hover:bg-accent hover:text-accent-foreground',
		outline: 'border border-input bg-background hover:bg-accent hover:text-accent-foreground'
	};

	const sizes: Record<string, string> = {
		default: 'h-10 px-4 py-2',
		sm: 'h-9 rounded-md px-3',
		icon: 'h-10 w-10'
	};
</script>

<button
	class={cn(
		'inline-flex items-center justify-center rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50',
		variants[variant],
		sizes[size],
		className
	)}
	{...rest}
>
	{#if children}{@render children()}{/if}
</button>
