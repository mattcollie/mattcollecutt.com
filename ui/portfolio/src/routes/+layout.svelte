<script lang="ts">
	import '../app.css';
	import { onMount } from 'svelte';
	import { browser } from '$app/environment';
	import * as publicEnv from '$env/static/public';

	let { children } = $props();

	onMount(() => {
		// namespace import: a missing var is undefined (analytics off), not a build failure
		if (browser && publicEnv.PUBLIC_POSTHOG_KEY) {
			import('posthog-js').then(({ default: posthog }) => {
				posthog.init(publicEnv.PUBLIC_POSTHOG_KEY, {
					api_host: publicEnv.PUBLIC_POSTHOG_HOST,
					person_profiles: 'always'
				});
			});
		}
	});
</script>

<svelte:head>
	<title>Matt Collecutt — Software Developer</title>
</svelte:head>

<div class="min-h-full bg-paper text-ink font-serif text-[17px] leading-[1.65]">
	{@render children()}
</div>
