<script>
    import "/src/app.css";

    import {onMount} from "svelte";
    import {fade, fly} from 'svelte/transition';

    import Analytics from "$lib/components/Analytics.svelte";
    import ThemeToggle from "$lib/components/ThemeToggle.svelte";
    import Link from "$lib/components/Link.svelte";

    let loading = true;
    onMount(() => {
        loading = false;
    });
</script>

<Analytics/>

<div class="min-h-full flex flex-col items-start pb-4">
    <ThemeToggle class="h-8 w-8 cursor-pointer mt-4 ml-4 lg:mt-8 lg:ml-8" />
    <div class="h-full w-full flex flex-col lg:flex-row gap-12 px-4 lg:py-36 lg:px-72">
        {#if !loading}
            <nav transition:fade class="flex flex-row lg:flex-col gap-1 pt-16 lg:pt-24">
                <Link href="/">home</Link>
                <Link href="/about">about</Link>
                <Link href="/blog">blog</Link>
                <Link href="/photography">photography</Link>
            </nav>
        {/if}

        {#if !loading}
            <div transition:fly={{y: 50, duration: 500}} class="h-full w-full flex flex-col">
                <slot/>
            </div>
        {/if}
    </div>
</div>