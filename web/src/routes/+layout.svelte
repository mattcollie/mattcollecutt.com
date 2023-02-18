<script>
    import "/src/app.css";
    import ThemeToggle from "$lib/ThemeToggle.svelte";
    import Link from "$lib/Link.svelte";
</script>

<div class="min-h-full flex flex-col items-start">
    <ThemeToggle class="h-8 cursor-pointer" />
    <div class="h-full flex flex-col lg:flex-row gap-12 lg:py-36 lg:px-96">
        <nav class="flex flex-row lg:flex-col gap-8 pt-16 lg:pt-24">
            <Link href="/">home</Link>
            <Link href="/about">about</Link>
            <Link href="/blog">blog</Link>
        </nav>

        <div class="h-full flex flex-col">
            <slot/>
        </div>
    </div>
</div>