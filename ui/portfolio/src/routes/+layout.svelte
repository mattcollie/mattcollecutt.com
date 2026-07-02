<script lang="ts">
	import '../app.css';
	import { onMount } from 'svelte';
	import { browser } from '$app/environment';
	import { PUBLIC_POSTHOG_KEY, PUBLIC_POSTHOG_HOST } from '$env/static/public';

	let { children } = $props();

	type Theme = 'auto' | 'light' | 'dark';
	let theme = $state<Theme>('auto');

	onMount(() => {
		const stored = localStorage.getItem('theme');
		if (stored === 'light' || stored === 'dark') theme = stored;

		if (browser && PUBLIC_POSTHOG_KEY) {
			import('posthog-js').then(({ default: posthog }) => {
				posthog.init(PUBLIC_POSTHOG_KEY, {
					api_host: PUBLIC_POSTHOG_HOST,
					person_profiles: 'always'
				});
			});
		}
	});

	function cycleTheme() {
		theme = theme === 'auto' ? 'light' : theme === 'light' ? 'dark' : 'auto';
		if (theme === 'auto') {
			localStorage.removeItem('theme');
			delete document.documentElement.dataset.theme;
		} else {
			localStorage.setItem('theme', theme);
			document.documentElement.dataset.theme = theme;
		}
	}
</script>

<svelte:head>
	<title>Matt Collecutt — Software Developer</title>
</svelte:head>

<div class="min-h-full bg-paper text-ink font-serif text-[17px] leading-[1.65]">
	<div class="max-w-[62ch] mx-auto px-6 pt-14 pb-24">
		<div class="flex items-baseline justify-between border-b border-rule pb-3.5 mb-14 font-mono text-[11.5px] text-grey">
			<span>mattcollecutt.com</span>
			<span class="flex items-baseline gap-4">
				<span class="hidden sm:inline">Waikato, New Zealand</span>
				<button
					type="button"
					class="cursor-pointer hover:text-accent transition-colors"
					onclick={cycleTheme}
					aria-label="Cycle theme (auto, light, dark)"
				>
					theme: {theme}
				</button>
			</span>
		</div>

		{@render children()}
	</div>
</div>
