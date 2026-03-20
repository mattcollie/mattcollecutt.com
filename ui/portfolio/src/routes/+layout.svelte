<script lang="ts">
	import '../app.css';

	let { children } = $props();
</script>

<svelte:head>
	<title>Matt Collecutt</title>
</svelte:head>

<div class="min-h-full bg-primary text-foreground font-sans">
	<header class="sticky top-0 z-10 bg-primary/80 backdrop-blur-sm border-b border-stroke">
		<div class="max-w-3xl mx-auto px-6 py-4 flex items-center justify-between">
			<a href="/" class="text-lg font-semibold text-black" style="font-family: 'Playfair Display', serif">Matt Collecutt</a>
			<div class="flex items-center gap-4">
				<a href="https://github.com/mattcollie" target="_blank" rel="noopener noreferrer" aria-label="GitHub">
					<img src="/images/github.svg" alt="GitHub" class="w-5 h-5 opacity-60 hover:opacity-100 transition-opacity" />
				</a>
				<a href="https://www.linkedin.com/in/mattcollecutt/" target="_blank" rel="noopener noreferrer" aria-label="LinkedIn">
					<img src="/images/linkedin.svg" alt="LinkedIn" class="w-5 h-5 opacity-60 hover:opacity-100 transition-opacity" />
				</a>
				<a href="https://www.instagram.com/mattcollecutt/" target="_blank" rel="noopener noreferrer" aria-label="Instagram">
					<img src="/images/instagram.svg" alt="Instagram" class="w-5 h-5 opacity-60 hover:opacity-100 transition-opacity" />
				</a>
			</div>
		</div>
	</header>

	{@render children()}
</div>
