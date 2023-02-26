<script>
    import { page } from '$app/stores';
</script>

<svelte:head>
    <title>{$page.status}</title>
</svelte:head>

<h1>{$page.status}</h1>

<div class="pt-12 flex flex-col gap-4">
    <p class="text-grey">{$page.error.message}</p>
</div>