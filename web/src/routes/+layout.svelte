<script>
    import "/src/app.css";

    import Analytics from "$lib/components/Analytics.svelte";
    import ThemeToggle from "$lib/components/ThemeToggle.svelte";
    import Link from "$lib/components/Link.svelte";
</script>

<Analytics/>

<div class="min-h-full flex flex-col items-start pb-4">
    <ThemeToggle class="h-8 w-8 cursor-pointer mt-4 ml-4 lg:mt-8 lg:ml-8" />
    <div class="h-full flex flex-col lg:flex-row gap-12 px-4 lg:py-36 lg:px-72">
        <nav class="flex flex-row lg:flex-col gap-1 pt-16 lg:pt-24">
            <Link href="/">home</Link>
            <Link href="/about">about</Link>
            <Link href="/blog">blog</Link>
            <Link href="/gallery">gallery</Link>
        </nav>

        <div class="h-full flex flex-col">
            <slot/>
        </div>
    </div>
</div>