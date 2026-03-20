import type { Snippet } from 'svelte';
import type { HTMLButtonAttributes } from 'svelte/elements';

export type ButtonVariant = 'default' | 'ghost' | 'outline';
export type ButtonSize = 'default' | 'sm' | 'icon';

export interface ButtonProps extends HTMLButtonAttributes {
	variant?: ButtonVariant;
	size?: ButtonSize;
	children?: Snippet;
}

export { default as Button } from './Button.svelte';
