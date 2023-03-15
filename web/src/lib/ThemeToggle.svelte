<script>
    import {darkTheme} from "$lib/store.js";
    import {onMount} from "svelte";
    import Moon from "$lib/assets/moon.svg";
    import Sun from "$lib/assets/sun.svg";

    onMount(() => {
        $darkTheme = localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)
    })

    const toggle = function() {
        $darkTheme = !$darkTheme;

        if(localStorage !== undefined) {
            localStorage.theme = $darkTheme ? 'dark' : 'light';
            if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
                document.documentElement.classList.add('dark')
                localStorage.theme = 'dark';
            } else {
                document.documentElement.classList.remove('dark')
                localStorage.theme = 'light';
            }
        }
    }
</script>

{#if $darkTheme}
    <img rel="prefetch" {...$$props} on:click={toggle} src={Sun} alt="Dark theme Icon">
{:else}
    <img rel="prefetch" {...$$props} on:click={toggle} src={Moon} alt="Light theme Icon">
{/if}