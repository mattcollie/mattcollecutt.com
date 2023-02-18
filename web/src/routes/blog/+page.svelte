<script>
    import { page } from "$lib/store";
    $page = "/blog";
    export let data;
    const {posts} = data;
</script>


<h1>blog</h1>

<div class="flex flex-col gap-4 pt-12">
    {#if posts.length > 0}
        {#each posts as {title, name, date}}
            <div>
                <h2><a href={'/blog/' + name}>{title}</a></h2>
                <p class="text-sm text-grey">{date}</p>
            </div>
        {/each}
    {:else}
        <p class="text-grey">Sorry, Can't find any posts at the moment!</p>
    {/if}
</div>