<script>
    import {page} from "$lib/store";
    import Link from "$lib/components/Link.svelte";

    $page = "/gallery";
    export let data = {galleries: {}};
    const {galleries} = data;
    const galleryOrder = Object.keys(galleries).sort((a, b) => new Date(galleries[b].date) - new Date(galleries[a].date));
</script>

<svelte:head>
    <title>Gallery</title>
</svelte:head>

<h1>Gallery</h1>

<div class="flex flex-col lg:flex-row gap-4 pt-6">
    {#if galleryOrder.length > 0}
        {#each galleryOrder as key}
            <Link href={'/gallery/'+key.replaceAll(' ', '_')}>
                <div class="relative flex flex-col gap-2">
                    <div class="relative photo-container m-auto">
                        <img rel="prefetch" src={galleries[key].photos[0].default}
                             alt={galleries[key].photos[0].default}/>
                    </div>
                    <p class="text-md text-grey-600 hover:text-black dark:text-grey-200 dark:hover:text-white">{key}</p>
                </div>
            </Link>
        {/each}
    {:else}
        <p class="text-grey">Sorry, Can't find any photo galleries at the moment!</p>
    {/if}
</div>

<style>
    .photo-container {
        width: 140px;
        height: 140px;
        position: relative;
    }

    img {
        width: 128px;
        height: 128px;
        @apply
            absolute m-auto object-cover
            hover:transition hover:ease-in-out hover:scale-[1.02]
    }
    img:nth-child(2) {
        top: 8px;
        left: 8px;
    }
    img:nth-child(3) {
        top: 16px;
        left: 16px;
    }
</style>