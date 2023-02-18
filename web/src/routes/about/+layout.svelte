<script>
    import { page } from "$lib/store";
    $page = "/about";
</script>

<svelte:head>
    <title>About</title>
</svelte:head>

<h1>About</h1>

<div class="pt-12 flex flex-col gap-4">
    <slot/>
</div>