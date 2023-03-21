<script>
    import {crossfade, scale} from 'svelte/transition';

    import {page} from "$lib/store";

    export let data = {photos: []};

    const {photos} = data;
    const photoOrder = photos.sort((a, b) => new Date(b.date) - new Date(b.date));
    const [send, receive] = crossfade({
        duration: 200,
        fallback: scale
    });

    let selected, loading;

    $page = "/photography";

    const load = async image => {
        const timeout = setTimeout(() => loading = image, 100);

        const img = new Image();

        img.onload = () => {
            selected = image;
            clearTimeout(timeout);
            loading = null;
        };

        img.src = image.photo.default;
    };
</script>

<svelte:head>
    <title>Photography</title>
</svelte:head>

<h1>Photography</h1>

<div class="grid pt-12">
    {#each photoOrder as photo}
        {#if selected !== photo.name}
            <button
                    in:receive={{key:photo.name}}
                    out:send={{key:photo.name}}
                    on:click="{() => load(photo)}">
                {#if loading === photo}
                    Loading...
                {:else}
                    <img
                            rel="prefetch"
                            alt={photo.name}
                            src={photo.photo.default}
                    >
                {/if}
            </button>
        {/if}
    {/each}
</div>

{#if selected}
    {#await selected then d}
        <button class="photo" in:receive={{key:d.name}} out:send={{key:d.name}} on:click={() => selected = null}>
            <img rel="prefetch" alt={d.name} src={d.photo.default}>
        </button>
    {/await}
{/if}


<style>
    .grid {
        flex: 1;
        width: 100%;
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(14rem, 2fr));
        grid-gap: 8rem;
    }

    button {
        width: 100%;
        height: 100%;
        will-change: transform;

        @apply
        h-64 w-64
    }

    .photo, img {
        @apply
        fixed top-0 left-0 w-full h-full object-fill overflow-hidden
    }

    img {
        @apply
        object-cover cursor-pointer
    }
</style>