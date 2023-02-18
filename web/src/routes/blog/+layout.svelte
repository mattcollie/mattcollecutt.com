<script>
    import { page } from "$lib/store";
    $page = "/blog";
</script>

<svelte:head>
    <title>blog</title>
</svelte:head>

<h1>blog</h1>

<div class="pt-12 flex flex-col gap-4">
    <slot/>
</div>