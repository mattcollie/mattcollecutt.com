import type { Component } from 'svelte';

export type NoteMeta = {
	title: string;
	date: string; // ISO, e.g. 2026-07-05
	draft?: boolean;
};

export type Note = NoteMeta & {
	slug: string;
	content: Component;
};

// Every .md in src/content/notes becomes a note. Publishing is: write the
// file, set draft: false (or remove it). Drafts never render anywhere.
const modules = import.meta.glob('/src/content/notes/*.md', { eager: true }) as Record<
	string,
	{ metadata: NoteMeta; default: Component }
>;

export const notes: Note[] = Object.entries(modules)
	.map(([path, mod]) => ({
		...mod.metadata,
		slug: path.split('/').pop()!.replace(/\.md$/, ''),
		content: mod.default
	}))
	.filter((n) => !n.draft)
	.sort((a, b) => b.date.localeCompare(a.date));
