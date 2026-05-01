<script lang="ts">
	import '../app.css';
	import { onMount } from 'svelte';
	import { browser } from '$app/environment';
	import { PUBLIC_POSTHOG_KEY, PUBLIC_POSTHOG_HOST } from '$env/static/public';

	let { children } = $props();

	onMount(() => {
		if (browser && PUBLIC_POSTHOG_KEY) {
			import('posthog-js').then(({ default: posthog }) => {
				posthog.init(PUBLIC_POSTHOG_KEY, {
					api_host: PUBLIC_POSTHOG_HOST,
					person_profiles: 'always'
				});
			});
		}
	});
</script>

<svelte:head>
	<title>Matt Collecutt — Software Developer</title>
</svelte:head>

<div class="grain min-h-full bg-primary text-foreground font-sans">
	<header class="sticky top-0 z-10 bg-primary/80 backdrop-blur-sm border-b border-stroke">
		<div class="max-w-3xl mx-auto px-6 py-4 flex items-center justify-between">
			<a href="/" class="text-lg font-semibold text-black">Matt Collecutt</a>
			<div class="flex items-center gap-4">
				<a href="https://github.com/mattcollie" target="_blank" rel="noopener noreferrer" aria-label="GitHub">
					<img src="/images/github.svg" alt="GitHub" class="w-5 h-5 opacity-60 hover:opacity-100 transition-opacity" />
				</a>
				<a href="https://www.linkedin.com/in/mattcollecutt/" target="_blank" rel="noopener noreferrer" aria-label="LinkedIn">
					<img src="/images/linkedin.svg" alt="LinkedIn" class="w-5 h-5 opacity-60 hover:opacity-100 transition-opacity" />
				</a>
				<a href="https://www.instagram.com/matthewcollecutt/" target="_blank" rel="noopener noreferrer" aria-label="Instagram">
					<img src="/images/instagram.svg" alt="Instagram" class="w-5 h-5 opacity-60 hover:opacity-100 transition-opacity" />
				</a>
			</div>
		</div>
	</header>

	{@render children()}
</div>
