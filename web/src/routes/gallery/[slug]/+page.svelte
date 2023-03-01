<script>
    import { fade } from 'svelte/transition'
    import { page } from "$lib/store";
    $page = "/gallery";
    export let data;
    let current = 0;
    $: currentPhoto = data.photos[current].toString();

    const increment = function(n) {
        if(current + n < 0) {
            current = 0;
        }
        else if(current + n < data.photos.length) {
            current += n;
        }
    }
</script>

<svelte:head>
    <title>{data.name}</title>
</svelte:head>

<div class="absolute h-screen w-screen top-0 left-0 bg-white dark:bg-black">
    <div class="relative h-screen w-screen flex flex-row justify-around lg:px-8">
        <div class="my-auto">
            <svg on:click={() => increment(-1)} class:disabled-button={current === 0} class:active-button={current >= 0} width="41" height="77" viewBox="0 0 41 77" fill="currentColor" xmlns="http://www.w3.org/2000/svg">
                <path d="M39.7072 0.707092L0.707155 39.7071M0.707153 38.2929L38.7072 76.2929" stroke="currentColor" stroke-width="2"/>
            </svg>
        </div>
        <div class="flex flex-col m-auto justify-around text-center justify-items-center gap-8">
            <p>{currentPhoto.slice(currentPhoto.lastIndexOf('/')+1)}</p>
            <div class="photo-container relative">
                {#each data.photos as photo, i}
                    {#if current === i}
                        <img class="absolute h-full lg:left-[12.5%] object-cover transition fade delay-150" transition:fade src={data.photos[i]} alt={data.photos[i]} />
                    {/if}
                {/each}
            </div>
            <a href="/gallery" class="text-grey text-xl hover:font-bold cursor-pointer">close</a>
        </div>
        <div class="my-auto">
            <svg on:click={() => increment(1)} class:disabled-button={current === data.photos.length-1} class:active-button={current !== data.photos.length-1} width="41" height="78" viewBox="0 0 41 78" fill="currentColor" xmlns="http://www.w3.org/2000/svg">
                <path d="M1 0.878662L40 39.8787M40 38.4644L2 76.4644" stroke="currentColor" stroke-width="2"/>
            </svg>
        </div>
    </div>
</div>

<style>
    .photo-container {
        min-width: 65vw;
        min-height: 65vh;

        @apply
            min-w-[80vw] min-h-[50vh]
    }

    .active-button {
        @apply
            cursor-pointer hover:transition lg:hover:scale-[1.05] hover:delay-100 hover:ease-in-out
    }

    .disabled-button {
        @apply
            text-grey cursor-not-allowed hover:transition-none lg:hover:scale-100
    }

    svg {
        @apply
            scale-[0.2] lg:scale-100
    }
</style>
