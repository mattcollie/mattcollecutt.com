<script>
    import { page } from "$lib/store";
    let blogPosts = [];
    $page = "/blog";
</script>

<h1>blog</h1>

<div class="pt-12 flex flex-col gap-4">
    {#if blogPosts.length > 0}
        {#each blogPosts as {title, date}}
            <h3>{title}</h3>
            <p class="text-grey">{date}</p>
        {/each}
    {:else}
        <p class="text-grey">Sorry, Can't find any posts at the moment!</p>
    {/if}
</div>