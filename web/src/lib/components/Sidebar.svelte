<script>
    export let show = false;

    const toggleShow = function(e) {
        show = !show;
        console.log('show', show)
    }
</script>

<style>
    nav {
        @apply
            absolute -translate-x-full z-30 flex h-full max-h-screen min-h-screen
            w-3/4 flex-none transform flex-col overflow-y-auto border-r
            bg-white pb-10 transition duration-300 ease-in-out
            sm:w-1/2 sm:pb-0 md:w-56 lg:relative
            lg:z-auto lg:w-56 lg:translate-x-0 2xl:w-72 2xl:w-80
            border-r-stroke bg-secondary;
    }

    .burger {
        @apply
            absolute p-2 my-2 mx-4
            rounded-md bg-black
            flex items-center justify-center
            cursor-pointer
            lg:hidden
            hover:bg-stroke
    }
</style>

<nav class:relative={show} class:!translate-x-0={show}>
    <slot/>
</nav>

<button class="burger" on:click={() => show = !show} aria-expanded={show} aria-haspopup="true" aria-label="Nav toggle">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="white" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-primary">
        <line x1="3" y1="12" x2="21" y2="12"></line>
        <line x1="3" y1="6" x2="21" y2="6"></line>
        <line x1="3" y1="18" x2="21" y2="18"></line>
    </svg>
</button>