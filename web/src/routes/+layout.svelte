<script>
    import "/src/app.css";
    import ThemeToggle from "$lib/ThemeToggle.svelte";
    import Link from "$lib/Link.svelte";
</script>

<div class="flex flex-col items-start">
    <ThemeToggle class="h-8" />
    <div class="flex flex-row gap-12 py-36 px-64">
        <nav class="flex flex-col gap-8 pt-24">
            <Link href="/">home</Link>
            <Link href="/about">about</Link>
            <Link href="/blog">blog</Link>
        </nav>

        <slot/>
    </div>
</div>