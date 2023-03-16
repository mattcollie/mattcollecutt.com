<script>
    import {page} from "$lib/store";

    $page = "/blog";
    export let data;
    const {posts} = data;
</script>

<svelte:head>
    <title>Blog</title>
</svelte:head>

<h1>Blog</h1>

<div class="flex flex-col gap-4 pt-10">
    {#if posts.length > 0}
        {#each posts as {title, name, date}}
            <div>
                <h2><a class="text-lg text-grey-600 hover:text-black dark:text-grey-200 dark:hover:text-white"
                       href={'/blog/' + name}>{title}</a></h2>
                <p class="text-sm text-grey font-mono">{date}</p>
            </div>
        {/each}
    {:else}
        <p class="text-grey">Sorry, Can't find any posts at the moment!</p>
    {/if}
</div>