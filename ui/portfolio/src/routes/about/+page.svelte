<script lang="ts">
	import PageHeader from '$lib/components/PageHeader.svelte';
</script>

<svelte:head>
	<title>About — Matt Collecutt</title>
</svelte:head>

<PageHeader />

<main class="max-w-[760px] mx-auto px-6 sm:px-10 pt-16 pb-28">
	<h1 class="font-sans font-medium text-[clamp(30px,4.5vw,42px)] leading-[1.15] tracking-[-0.01em] mb-8">
		About
	</h1>

	<p class="text-[19px] max-w-[60ch] mb-2">
		I'm a software developer. The longer answer is harder to distil — the work at
		Mustard encompasses a lot of things, payments and finance mostly. What I love is
		fighting fires quickly and calmly, and designing and building systems. Taniwha is
		younger — there we're looking at what AI might mean next, starting with cognitive
		agents and building worlds around them.
	</p>

	<h2 class="font-sans font-medium text-[19px] mt-[60px] mb-4">The record, for the curious</h2>
	<div class="border-t border-rule">
		<div class="flex justify-between gap-4 py-3.5 border-b border-rule">
			<div class="min-w-0">
				<div>
					<a class="doc-a font-medium no-underline hover:underline" href="https://www.mustard.co.nz/" target="_blank" rel="noopener noreferrer">Mustard Technology</a>
					<span class="text-grey text-[15px]">— Developer</span>
				</div>
				<div class="text-[15.5px] mt-0.5">Payments and finance systems, mostly — it encompasses a lot of things.</div>
			</div>
			<div class="font-mono text-xs text-grey whitespace-nowrap pt-[5px]">2018 — now</div>
		</div>
		<div class="flex justify-between gap-4 py-3.5 border-b border-rule">
			<div class="min-w-0">
				<div>
					<a class="doc-a font-medium no-underline hover:underline" href="https://taniwha.ai" target="_blank" rel="noopener noreferrer">Taniwha AI</a>
					<span class="text-grey text-[15px]">— Developer</span>
				</div>
				<div class="text-[15.5px] mt-0.5">A cognitive engine giving AI agents real beliefs, memory, trust, and identity.</div>
			</div>
			<div class="font-mono text-xs text-grey whitespace-nowrap pt-[5px]">2026 — now</div>
		</div>
		<div class="flex justify-between gap-4 py-3.5 border-b border-rule">
			<div class="min-w-0">
				<span class="font-medium">Ezitracker ANZ</span>
				<span class="text-grey text-[15px]">— Junior Developer</span>
			</div>
			<div class="font-mono text-xs text-grey whitespace-nowrap pt-[5px]">2017 — 2018</div>
		</div>
		<div class="flex justify-between gap-4 py-3.5 border-b border-rule">
			<div class="min-w-0">
				<span class="font-medium">Panztel</span>
				<span class="text-grey text-[15px]">— Apprentice Developer</span>
			</div>
			<div class="font-mono text-xs text-grey whitespace-nowrap pt-[5px]">2015 — 2017</div>
		</div>
	</div>

	<h2 class="font-sans font-medium text-[19px] mt-[60px] mb-4">This site</h2>
	<p class="text-[17px] max-w-[62ch]">
		Built with SvelteKit, set in Jost and Newsreader. The front page draws the real
		contour lines of two hills that matter to me, from
		<a class="doc-a" href="https://data.linz.govt.nz/layer/50768-nz-contours-topo-150k/" target="_blank" rel="noopener noreferrer">LINZ Topo50 data</a>
		(CC BY 4.0). Claude does the plumbing; the words and photographs are mine.
	</p>

	<p class="mt-16 pt-4 border-t border-rule font-mono text-[11.5px]">
		<a class="doc-a" href="/">← mattcollecutt.com</a>
	</p>
</main>
