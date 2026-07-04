<script lang="ts">
	import { onMount } from 'svelte';
	import ThemeToggle from '$lib/components/ThemeToggle.svelte';

	type MapKey = 'teAroha' | 'redHill';
	type ContourMap = { w: number; h: number; levels: { lv: number; lines: number[][][] }[] };

	const captions: Record<MapKey, string> = {
		teAroha: 'Te Aroha & the Kaimai Range — the mountain I grew up under.',
		redHill: 'Red Hill & the Hunua Ranges — the hill I live on now.'
	};

	let mapKey = $state<MapKey>('teAroha');
	let maps: Record<MapKey, ContourMap> | null = null;
	let canvas: HTMLCanvasElement;

	const REVEAL_MS = 3200; // plains first, summit last
	let raf = 0;
	let start: number | null = null;
	let reduced = false;

	function mapColor() {
		return getComputedStyle(document.documentElement).getPropertyValue('--map').trim();
	}

	function draw(ts: number, instant = false) {
		if (!maps) return;
		const data = maps[mapKey];
		if (start === null) start = ts;
		const t = instant || reduced ? 1 : Math.min((ts - start) / REVEAL_MS, 1);
		const ctx = canvas.getContext('2d');
		if (!ctx) return;
		const dpr = devicePixelRatio || 1;
		const { width: W, height: H } = canvas;
		ctx.clearRect(0, 0, W, H);
		// cover-fit the map frame into the hero
		const s = Math.max(W / data.w, H / data.h);
		const ox = (W - data.w * s) / 2;
		const oy = (H - data.h * s) / 2;
		ctx.lineWidth = 1 * dpr;
		const rgb = mapColor();
		const n = data.levels.length;
		for (let i = 0; i < n; i++) {
			// each elevation fades in during its slice of the reveal
			const lt = Math.max(0, Math.min((t * n - i) / 1.6, 1));
			if (lt <= 0) continue;
			ctx.strokeStyle = `rgba(${rgb},${(0.34 - i * 0.006) * lt})`;
			for (const ln of data.levels[i].lines) {
				ctx.beginPath();
				for (let k = 0; k < ln.length; k++) {
					const x = ox + ln[k][0] * s;
					const y = oy + ln[k][1] * s;
					if (k === 0) ctx.moveTo(x, y);
					else ctx.lineTo(x, y);
				}
				ctx.stroke();
			}
		}
		if (t < 1) raf = requestAnimationFrame((ts2) => draw(ts2));
	}

	function restart(instant = false) {
		cancelAnimationFrame(raf);
		const dpr = devicePixelRatio || 1;
		canvas.width = canvas.clientWidth * dpr;
		canvas.height = canvas.clientHeight * dpr;
		start = null;
		raf = requestAnimationFrame((ts) => draw(ts, instant));
	}

	function selectMap(k: MapKey) {
		mapKey = k;
		restart();
	}

	onMount(() => {
		reduced = matchMedia('(prefers-reduced-motion: reduce)').matches;
		const redraw = () => restart(true);
		const mq = matchMedia('(prefers-color-scheme: dark)');
		window.addEventListener('themechange', redraw);
		mq.addEventListener('change', redraw);

		fetch('/map/contours.json')
			.then((r) => r.json())
			.then((data) => {
				maps = data;
				mapKey = Math.random() < 0.5 ? 'teAroha' : 'redHill';
				restart();
			});

		return () => {
			cancelAnimationFrame(raf);
			window.removeEventListener('themechange', redraw);
			mq.removeEventListener('change', redraw);
		};
	});
</script>

<svelte:window onresize={() => maps && restart(true)} />

<div class="relative min-h-screen">
	<canvas bind:this={canvas} class="absolute inset-0 w-full h-full" aria-hidden="true"></canvas>

	<div class="relative flex justify-between px-6 sm:px-10 py-4 font-mono text-[11px] text-grey">
		<span>mattcollecutt.com</span>
		<ThemeToggle />
	</div>

	<div class="relative flex flex-col justify-center px-6 sm:px-10 max-w-[760px] min-h-[calc(100vh-120px)]">
		<h1 class="font-sans font-medium text-[clamp(44px,6.5vw,78px)] tracking-[-0.02em] leading-[1.05] mb-3.5">
			Matt Collecutt
		</h1>
		<p class="text-[21px] max-w-[46ch]">
			I write software at
			<a class="doc-a" href="https://www.mustard.co.nz/" target="_blank" rel="noopener noreferrer">Mustard</a>
			and
			<a class="doc-a" href="https://taniwha.ai" target="_blank" rel="noopener noreferrer">Taniwha AI</a>.
			<span class="placeholder-note">(a sentence or two of yours — what you're making and thinking about lately)</span>
		</p>
	</div>

	<div class="absolute right-4 bottom-3.5 max-w-[46ch] text-right font-mono text-[10.5px] leading-[1.7] text-grey">
		<span>{captions[mapKey]}</span><br />
		<button type="button" class="map-btn" aria-pressed={mapKey === 'teAroha'} onclick={() => selectMap('teAroha')}>te aroha</button>
		·
		<button type="button" class="map-btn" aria-pressed={mapKey === 'redHill'} onclick={() => selectMap('redHill')}>red hill</button><br />
		<span class="opacity-75">
			contours from the
			<a class="text-grey underline underline-offset-[3px]" href="https://data.linz.govt.nz/layer/50768-nz-contours-topo-150k/" target="_blank" rel="noopener noreferrer">LINZ Data Service</a>
			(Topo50), CC BY 4.0 · how this map is made — <span class="italic">note coming</span>
		</span>
	</div>
</div>

<main class="max-w-[760px] mx-auto px-6 sm:px-10 pt-10">
	<h2 class="font-sans font-medium text-[19px] mt-[70px] mb-4">Notes</h2>
	<div class="flex gap-4 items-baseline py-3 border-t border-rule">
		<span class="font-mono text-xs text-grey shrink-0">····</span>
		<span class="text-[21px] placeholder-note">How I actually use AI to build software (draft in progress)</span>
	</div>
	<div class="flex gap-4 items-baseline py-3 border-t border-b border-rule">
		<span class="font-mono text-xs text-grey shrink-0">····</span>
		<span class="text-[21px] placeholder-note">First vintage: winemaking notes from the shed (example title)</span>
	</div>

	<h2 class="font-sans font-medium text-[19px] mt-[70px] mb-4">Photographs</h2>
	<div class="grid grid-cols-1 sm:grid-cols-[1.5fr_1fr] gap-3">
		<div class="rounded-[3px] min-h-[320px] bg-[radial-gradient(130%_100%_at_20%_0%,#2E5D49,#1F4A38_55%,#10241B)]"></div>
		<div class="rounded-[3px] min-h-[320px] bg-[radial-gradient(120%_110%_at_80%_15%,#526030,#333F1C_55%,#151B0B)]"></div>
	</div>
	<p class="font-mono text-[11px] text-grey mt-2 placeholder-note">(two of yours — captions say where and when)</p>

	<p class="mt-20 pt-4 border-t border-rule text-[16.5px] text-grey">
		Small things I've built for myself over the years, mostly older and no longer maintained:
		<a class="text-grey underline underline-offset-[3px]" href="https://trivino.xyz" target="_blank" rel="noopener noreferrer">TriVino</a>,
		<a class="text-grey underline underline-offset-[3px]" href="https://nookly.co.nz" target="_blank" rel="noopener noreferrer">Nookly</a>,
		<a class="text-grey underline underline-offset-[3px]" href="https://brewlog.co.nz" target="_blank" rel="noopener noreferrer">Brewlog</a>.
	</p>
</main>

<footer class="max-w-[760px] mx-auto px-6 sm:px-10 pt-12 pb-24 flex justify-between gap-3 flex-wrap font-mono text-[11.5px] text-grey">
	<span>
		&copy; {new Date().getFullYear()} Matt Collecutt ·
		<a class="doc-a" href="https://github.com/mattcollie" target="_blank" rel="noopener noreferrer">github</a> ·
		<a class="doc-a" href="https://www.linkedin.com/in/mattcollecutt/" target="_blank" rel="noopener noreferrer">linkedin</a> ·
		<a class="doc-a" href="https://www.instagram.com/matthewcollecutt/" target="_blank" rel="noopener noreferrer">instagram</a> ·
		<a class="doc-a" href="mailto:matt@mustard.co.nz">email</a>
	</span>
	<span class="placeholder-note">(about page coming)</span>
</footer>

<style>
	.map-btn {
		all: unset;
		cursor: pointer;
		font-family: var(--font-mono);
		font-size: 10.5px;
		color: var(--grey);
		text-decoration: underline;
		text-underline-offset: 3px;
	}
	.map-btn[aria-pressed='true'] {
		color: var(--accent);
		text-decoration: none;
	}
	.map-btn:focus-visible {
		outline: 2px solid var(--accent);
	}
</style>
