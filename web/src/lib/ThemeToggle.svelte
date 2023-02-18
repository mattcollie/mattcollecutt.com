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

<img on:click={toggle} src={$darkTheme ? Moon : Sun} alt="Current theme icon">