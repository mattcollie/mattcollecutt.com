<script>
    import { page } from "$lib/store";
    import Link from "$lib/Link.svelte";
    $page = "/gallery";
    export let data;
    const {galleries} = data;
    const galleryOrder = Object.keys(galleries).sort((a, b) => new Date(galleries[b].date) - new Date(galleries[a].date));
</script>

<svelte:head>
    <title>Gallery</title>
</svelte:head>

<h1>Gallery</h1>

<div class="flex flex-col lg:flex-row gap-4 pt-12">
    {#if galleryOrder.length > 0}
        {#each galleryOrder as key}
            <Link href={'/gallery/'+key.replaceAll(' ', '_')}>
                <div class="relative flex flex-col gap-4">
                    <div class="relative photo-container m-auto">
                        {#each galleries[key].photos.slice(0, 3).reverse() as photo}
                            <img src={photo.default} alt={photo.default} />
                        {/each}
                    </div>
                    <p class="text-md text-black dark:text-white text-center">{key}</p>
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