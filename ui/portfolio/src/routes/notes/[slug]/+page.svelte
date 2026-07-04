<script lang="ts">
	import PageHeader from '$lib/components/PageHeader.svelte';

	let { data } = $props();
	const NoteContent = $derived(data.note.content);
	const dateFmt = $derived(
		new Date(data.note.date + 'T00:00:00').toLocaleDateString('en-NZ', {
			year: 'numeric',
			month: 'long'
		})
	);
</script>

<svelte:head>
	<title>{data.note.title} — Matt Collecutt</title>
</svelte:head>

<PageHeader />

<article class="max-w-[760px] mx-auto px-6 sm:px-10 pt-16 pb-28">
	<p class="font-mono text-xs text-grey mb-3">{dateFmt}</p>
	<h1 class="font-sans font-medium text-[clamp(30px,4.5vw,42px)] leading-[1.15] tracking-[-0.01em] mb-10">
		{data.note.title}
	</h1>
	<div class="note-body">
		<NoteContent />
	</div>
	<p class="mt-16 pt-4 border-t border-rule font-mono text-[11.5px]">
		<a class="doc-a" href="/">← mattcollecutt.com</a>
	</p>
</article>

<style>
	.note-body :global(p) {
		margin: 0 0 1em;
		font-size: 18px;
		line-height: 1.7;
		max-width: 68ch;
	}
	.note-body :global(h2) {
		font-family: var(--font-sans);
		font-weight: 500;
		font-size: 21px;
		margin: 2.2em 0 0.7em;
	}
	.note-body :global(h3) {
		font-family: var(--font-sans);
		font-weight: 500;
		font-size: 17px;
		margin: 1.8em 0 0.6em;
	}
	.note-body :global(a) {
		color: var(--accent);
		text-decoration: underline;
		text-decoration-thickness: 1px;
		text-underline-offset: 3px;
	}
	.note-body :global(blockquote) {
		margin: 1.2em 0;
		padding-left: 1.1em;
		border-left: 2px solid var(--rule);
		color: var(--grey);
		font-style: italic;
	}
	.note-body :global(code) {
		font-family: var(--font-mono);
		font-size: 0.85em;
	}
	.note-body :global(pre) {
		font-family: var(--font-mono);
		font-size: 13px;
		line-height: 1.6;
		padding: 14px 16px;
		border: 1px solid var(--rule);
		border-radius: 3px;
		overflow-x: auto;
		margin: 1.2em 0;
	}
	.note-body :global(ul),
	.note-body :global(ol) {
		padding-left: 1.4em;
		margin: 0 0 1em;
		font-size: 18px;
		line-height: 1.7;
	}
	.note-body :global(li) {
		margin-bottom: 0.3em;
	}
	.note-body :global(hr) {
		border: 0;
		border-top: 1px solid var(--rule);
		margin: 2em 0;
	}
	.note-body :global(img) {
		max-width: 100%;
		border-radius: 3px;
	}
	.note-body :global(em) {
		font-style: italic;
	}
</style>
