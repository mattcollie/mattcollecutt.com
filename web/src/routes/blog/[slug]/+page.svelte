<script>
    import { page } from "$lib/store";
    $page = "/blog";
    export let data;
</script>

<svelte:head>
    <title>{data.title}</title>
</svelte:head>


<h1>{data.title}</h1>
<p class="text-sm text-grey">{data.date}</p>

<article class="flex flex-col gap-4 pt-8">
        <svelte:component this={data.content} />
</article>
