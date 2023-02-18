<script>
    let blogPosts = [];
</script>

{#if blogPosts.length > 0}
    {#each blogPosts as {title, date}}
        <h3>{title}</h3>
        <p class="text-grey">{date}</p>
    {/each}
{:else}
    <p class="text-grey">Sorry, Can't find any posts at the moment!</p>
{/if}