<script lang="ts">
	import '../app.css';
	import Sidebar from '$lib/components/Sidebar.svelte';
	import Banner from '$lib/components/Banner.svelte';

	let { children } = $props();
	let mobileOpen = $state(false);
</script>

<svelte:head>
	<title>Matt Collecutt</title>
</svelte:head>

<div class="min-h-full bg-black text-white font-sans">
	<div class="relative h-screen flex flex-row">
		<Sidebar />
		<Sidebar mobile={true} bind:open={mobileOpen} />

		<div class="h-full w-full flex flex-col lg:ml-56 2xl:ml-80">
			<!-- Mobile hamburger -->
			<div class="lg:hidden">
				<button
					class="absolute p-2 my-2 mx-4 rounded-md bg-black flex items-center justify-center cursor-pointer hover:bg-stroke z-20"
					onclick={() => mobileOpen = true}
					aria-label="Open navigation"
					aria-haspopup="true"
				>
					<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="white" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
						<line x1="3" y1="12" x2="21" y2="12"></line>
						<line x1="3" y1="6" x2="21" y2="6"></line>
						<line x1="3" y1="18" x2="21" y2="18"></line>
					</svg>
				</button>
			</div>

			{@render children()}
		</div>
	</div>
</div>
