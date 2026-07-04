<script lang="ts">
	import { photos } from '$lib/photos';

	let current = $state(-1);
	let dialogEl: HTMLDialogElement;

	function show(i: number) {
		current = i;
		dialogEl.showModal();
	}

	function close() {
		dialogEl.close();
		current = -1;
	}

	function nav(d: number) {
		current = (current + d + photos.length) % photos.length;
	}

	function onKey(e: KeyboardEvent) {
		if (current < 0) return;
		if (e.key === 'ArrowRight') nav(1);
		if (e.key === 'ArrowLeft') nav(-1);
	}

	function onDialogClick(e: MouseEvent) {
		// a click on the backdrop (the dialog element itself) closes
		if (e.target === dialogEl) close();
	}
</script>

<svelte:window onkeydown={onKey} />

<!-- The mosaic (mixed spans, cropped tiles) needs mass; below three photos,
     show every frame uncropped side by side. -->
{#if photos.length < 3}
	<div class="grid grid-cols-1 sm:grid-cols-2 gap-3">
		{#each photos as photo, i}
			<figure class="m-0">
				<button
					type="button"
					class="block w-full cursor-zoom-in focus-visible:outline-2 focus-visible:outline-accent"
					onclick={() => show(i)}
					aria-label="View larger: {photo.caption}"
				>
					<img src={photo.src} alt={photo.alt} decoding="async" class="rounded-[3px] w-full h-auto" />
				</button>
				<figcaption class="font-mono text-[11px] text-grey mt-2">{photo.caption}</figcaption>
			</figure>
		{/each}
	</div>
{:else}
	<div class="grid grid-cols-1 sm:grid-cols-3 gap-3">
		{#each photos as photo, i}
			<figure class="m-0 {photo.size === 'wide' ? 'sm:col-span-2' : ''}">
				<button
					type="button"
					class="block w-full cursor-zoom-in focus-visible:outline-2 focus-visible:outline-accent"
					onclick={() => show(i)}
					aria-label="View larger: {photo.caption}"
				>
					<img
						src={photo.src}
						alt={photo.alt}
						decoding="async"
						class="rounded-[3px] w-full h-60 sm:h-80 lg:h-96 object-cover"
					/>
				</button>
				<figcaption class="font-mono text-[11px] text-grey mt-2">{photo.caption}</figcaption>
			</figure>
		{/each}
	</div>
{/if}

<dialog
	bind:this={dialogEl}
	onclick={onDialogClick}
	onclose={() => (current = -1)}
	class="m-auto bg-transparent p-0 max-w-none border-0 outline-none backdrop:bg-[color-mix(in_srgb,var(--paper)_92%,transparent)]"
>
	{#if current >= 0}
		<figure class="m-0 flex flex-col items-center gap-3 p-4">
			<img
				src={photos[current].full}
				alt={photos[current].alt}
				class="max-h-[82vh] max-w-[92vw] w-auto h-auto rounded-[3px]"
			/>
			<figcaption class="font-mono text-[11.5px] text-grey flex items-center gap-5">
				{#if photos.length > 1}
					<button type="button" class="lb-btn" onclick={() => nav(-1)} aria-label="Previous photograph">←</button>
				{/if}
				<span class="text-ink">{photos[current].caption}</span>
				{#if photos.length > 1}
					<button type="button" class="lb-btn" onclick={() => nav(1)} aria-label="Next photograph">→</button>
				{/if}
				<button type="button" class="lb-btn" onclick={close} aria-label="Close">esc</button>
			</figcaption>
		</figure>
	{/if}
</dialog>

<style>
	.lb-btn {
		all: unset;
		cursor: pointer;
		font-family: var(--font-mono);
		font-size: 11.5px;
		color: var(--grey);
		text-decoration: underline;
		text-underline-offset: 3px;
	}
	.lb-btn:hover {
		color: var(--accent);
	}
	.lb-btn:focus-visible {
		outline: 2px solid var(--accent);
	}
</style>
