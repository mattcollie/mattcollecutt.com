<script lang="ts">
	import { onMount } from 'svelte';

	type Theme = 'auto' | 'light' | 'dark';
	let theme = $state<Theme>('auto');

	onMount(() => {
		const stored = localStorage.getItem('theme');
		if (stored === 'light' || stored === 'dark') theme = stored;
	});

	function cycle() {
		theme = theme === 'auto' ? 'light' : theme === 'light' ? 'dark' : 'auto';
		if (theme === 'auto') {
			localStorage.removeItem('theme');
			delete document.documentElement.dataset.theme;
		} else {
			localStorage.setItem('theme', theme);
			document.documentElement.dataset.theme = theme;
		}
		window.dispatchEvent(new CustomEvent('themechange'));
	}
</script>

<button
	type="button"
	class="cursor-pointer hover:text-accent transition-colors"
	onclick={cycle}
	aria-label="Cycle theme (auto, light, dark)"
>
	theme: {theme}
</button>
