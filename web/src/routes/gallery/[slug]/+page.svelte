<script>
    import {Swipeable} from 'thumb-ui'
    import { fade } from 'svelte/transition'
    import { page } from "$lib/store";
    $page = "/gallery";
    export let data;

    let current, loginProgress, introProgress;
</script>

<style>
    .fullpage {
        position: absolute;
        width: 100%;
        height: 100%;
    }
    .slides {
        text-align: center;
    }
    .current {
        transform: scale(1);
    }

    .dots {
        position: absolute;
        bottom: 5%;
        left: 50%;
        transform: translateX(-50%);
        line-height: 0;

        @apply
            flex flex-row gap-2
    }
</style>

<svelte:head>
    <title>{data.name}</title>
</svelte:head>

<div class="absolute bg-white dark:bg-black top-0 left-0 h-full w-full m-0">
    <a class="absolute scale-50 cursor-pointer z-50 hover:scale-[0.60] transition top-5 left-5" href="/gallery">
        <svg width="59" height="60" viewBox="0 0 59 60" fill="currentColor" xmlns="http://www.w3.org/2000/svg">
            <path d="M1.35352 1.64645L57.9221 58.215M1.35352 58.9221L57.9221 2.35356" stroke="currentColor" stroke-width="3"/>
        </svg>
    </a>
    <div class="slides fullpage" style="transform: scale({1 - $loginProgress*0.3})">
        <Swipeable numScreens={data.photos.length} bind:current bind:progress={introProgress}>
            {#each data.photos as photo, i}
                <section class="absolute left-0 right-0 top-0 bottom-0 lg:mx-auto w-[100vw] h-full lg:w-[95vw] p-2 lg:p-10" class:current={current === i}>
                    <div class="h-full flex flex-col justify-around lg:justify-between py-[8vh]" style="opacity: {1 - Math.abs($introProgress - i)}">
                        <p>{photo.slice(photo.lastIndexOf('/')+1)}</p>
                        <img class="h-[80vh] lg:h-[70vh] lg:w-[90vw] object-cover lg:object-contain" src={photo} alt={photo}/>
                    </div>
                </section>
            {/each}
            <div class="dots flex flex-row gap-2">
                {#each data.photos as p, i}
                    <div class="w-3 h-3 bg-grey rounded-xl {current === i ? 'border-black dark:border-white border-2' : ''}" in:fade on:click={() => {current = i; $introProgress = i}}></div>
                {/each}
            </div>
        </Swipeable>
    </div>
</div>
